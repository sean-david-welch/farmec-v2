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
    endpoint: URL;
    queryKey: string;
}

export interface Resources {
    [key: string]: ResourceConfig;
}
