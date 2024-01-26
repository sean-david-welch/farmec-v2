export interface Machine {
    id?: string;
    supplierId: string;
    name: string;
    machine_image: string;
    description: string;
    machine_link?: string;
    created?: string;
}

export interface Product {
    id?: string;
    machineId: string;
    name: string;
    product_image: string;
    description: string;
    product_link: string;
}

export interface Sparepart {
    id?: string;
    supplierId: string;
    name: string;
    parts_image: string;
    spare_parts_link: string;
}

export interface Supplier {
    id?: string;
    name: string;
    logo_image: string;
    marketing_image: string;
    description: string;
    social_facebook?: string;
    social_instagram?: string;
    social_linkedin?: string;
    social_twitter?: string;
    social_youtube?: string;
    social_website?: string;
    created?: string;
}
