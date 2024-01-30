export interface SupplierResult {
    presignedLogoUrl: string;
    logoUrl: string;
    presignedMarketingUrl: string;
    marketingUrl: string;
}

export interface ModelResult {
    presignedUrl: string;
    imageUrl: string;
}

export interface EmailData {
    name: string;
    email: string;
    message: string;
}

export interface UserData {
    email: string;
    password: string;
    role: string;
}

export interface ResourceData {
    id: string;
    route: string;
    queryKey: string;
}

export interface ResourceConfig {
    endpoint: string;
    queryKey: string;
}

export interface Resources {
    [key: string]: ResourceConfig;
}

export interface FormField {
    name: string;
    label: string;
    type: string;
    options?: { label: string; value: string | undefined }[];
    placeholder: string;
    defaultValue?: string;
}
