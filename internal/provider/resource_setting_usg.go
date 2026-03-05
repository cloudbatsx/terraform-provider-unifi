package provider

import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudbatsx/go-unifi/unifi"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var resourceSettingUsgLock = sync.Mutex{}

func resourceSettingUsgLocker(f func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics) func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	return func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		resourceSettingUsgLock.Lock()
		defer resourceSettingUsgLock.Unlock()
		return f(ctx, d, meta)
	}
}

func resourceSettingUsg() *schema.Resource {
	return &schema.Resource{
		Description: "`unifi_setting_usg` manages settings for a Unifi Security Gateway.",

		CreateContext: resourceSettingUsgLocker(resourceSettingUsgUpsert),
		ReadContext:   resourceSettingUsgLocker(resourceSettingUsgRead),
		UpdateContext: resourceSettingUsgLocker(resourceSettingUsgUpsert),
		DeleteContext: schema.NoopContext,
		Importer: &schema.ResourceImporter{
			StateContext: importSiteAndID,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the settings.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"site": {
				Description: "The name of the site to associate the settings with.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
			},
			"multicast_dns_enabled": {
				Description: "Whether multicast DNS is enabled.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSettingUsgUpdateResourceData(d *schema.ResourceData, meta interface{}, setting *unifi.SettingUsg) error {
	c := meta.(*client)

	//nolint // GetOkExists is deprecated, but using here:
	if mdns, hasMdns := d.GetOkExists("multicast_dns_enabled"); hasMdns {
		if v := c.ControllerVersion(); v.GreaterThanOrEqual(controllerV7) {
			return fmt.Errorf("multicast_dns_enabled is not supported on controller version %v", c.ControllerVersion())
		}

		setting.MdnsEnabled = mdns.(bool)
	}

	return nil
}

func resourceSettingUsgUpsert(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client)

	site := d.Get("site").(string)
	if site == "" {
		site = c.site
	}

	req, err := c.c.GetSettingUsg(ctx, c.site)
	if err != nil {
		return diag.FromErr(err)
	}

	err = resourceSettingUsgUpdateResourceData(d, meta, req)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, err := c.c.UpdateSettingUsg(ctx, site, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.ID)
	return resourceSettingUsgSetResourceData(resp, d, meta, site)
}

func resourceSettingUsgSetResourceData(resp *unifi.SettingUsg, d *schema.ResourceData, meta interface{}, site string) diag.Diagnostics {
	d.Set("site", site)
	d.Set("multicast_dns_enabled", resp.MdnsEnabled)

	return nil
}

func resourceSettingUsgRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client)

	site := d.Get("site").(string)
	if site == "" {
		site = c.site
	}

	resp, err := c.c.GetSettingUsg(ctx, site)
	if _, ok := err.(*unifi.NotFoundError); ok {
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSettingUsgSetResourceData(resp, d, meta, site)
}
