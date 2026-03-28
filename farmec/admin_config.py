from django.urls import reverse_lazy

UNFOLD = {
    "SITE_TITLE": "Farmec Admin",
    "SITE_HEADER": "Farmec Administration",
    "SITE_SYMBOL": "agriculture",
    "SITE_FAVICON": lambda request: "/static/favicon.svg",
    "COLORS": {
        "primary": {
            "50": "254 242 242",
            "100": "254 226 226",
            "200": "254 202 202",
            "300": "252 165 165",
            "400": "248 113 113",
            "500": "220 38 38",
            "600": "185 28 28",
            "700": "153 27 27",
            "800": "127 29 29",
            "900": "127 29 29",
            "950": "69 10 10",
        },
        "gray": {
            "50": "250 250 250",
            "100": "245 245 245",
            "200": "229 229 229",
            "300": "212 212 212",
            "400": "163 163 163",
            "500": "115 115 115",
            "600": "82 82 82",
            "700": "64 64 64",
            "800": "38 38 38",
            "900": "10 10 10",
            "950": "0 0 0",
        },
    },
    "SIDEBAR": {
        "show_search": True,
        "show_all_applications": False,
        "navigation": [
            {
                "title": "Catalog",
                "separator": True,
                "items": [
                    {"title": "Suppliers", "icon": "factory", "link": reverse_lazy("admin:catalog_supplier_changelist")},
                    {"title": "Machines", "icon": "agriculture", "link": reverse_lazy("admin:catalog_machine_changelist")},
                    {"title": "Products", "icon": "inventory_2", "link": reverse_lazy("admin:catalog_product_changelist")},
                    {"title": "Spare Parts", "icon": "settings_suggest", "link": reverse_lazy("admin:catalog_spareparts_changelist")},
                    {"title": "Videos", "icon": "play_circle", "link": reverse_lazy("admin:catalog_video_changelist")},
                ],
            },
            {
                "title": "Content",
                "separator": True,
                "items": [
                    {"title": "Blog Posts", "icon": "article", "link": reverse_lazy("admin:content_blog_changelist")},
                    {"title": "Carousel", "icon": "view_carousel", "link": reverse_lazy("admin:content_carousel_changelist")},
                    {"title": "Exhibitions", "icon": "event", "link": reverse_lazy("admin:content_exhibition_changelist")},
                    {"title": "Timeline", "icon": "timeline", "link": reverse_lazy("admin:content_timeline_changelist")},
                ],
            },
            {
                "title": "Team",
                "separator": True,
                "items": [
                    {"title": "Employees", "icon": "badge", "link": reverse_lazy("admin:team_employee_changelist")},
                ],
            },
            {
                "title": "Support",
                "separator": True,
                "items": [
                    {"title": "Warranty Claims", "icon": "gavel", "link": reverse_lazy("admin:support_warrantyclaim_changelist")},
                    {"title": "Machine Registrations", "icon": "app_registration", "link": reverse_lazy("admin:support_machineregistration_changelist")},
                ],
            },
            {
                "title": "Legal",
                "separator": True,
                "items": [
                    {"title": "Privacy Policy", "icon": "policy", "link": reverse_lazy("admin:legal_privacy_changelist")},
                    {"title": "Terms & Conditions", "icon": "description", "link": reverse_lazy("admin:legal_terms_changelist")},
                ],
            },
        ],
    },
}
