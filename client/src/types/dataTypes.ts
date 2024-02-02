import { Employee, Privacy, Terms, Timeline } from './aboutTypes';
import Blog, { Exhibition } from './blogTypes';
import { Carousel, LineItem, MachineRegistration, WarrantyClaim } from './miscTypes';
import { Machine, Product, Sparepart, Supplier } from './supplierTypes';
import { Video } from './videoTypes';

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

export interface ResourceEntry<T> {
    endpoint: string;
    queryKey: string;
    type?: T;
}

export interface Resources {
    suppliers: ResourceEntry<Supplier>;
    spareparts: ResourceEntry<Sparepart>;
    machines: ResourceEntry<Machine>;
    products: ResourceEntry<Product>;
    videos: ResourceEntry<Video>;
    blogs: ResourceEntry<Blog>;
    exhibitions: ResourceEntry<Exhibition>;
    employees: ResourceEntry<Employee>;
    timelines: ResourceEntry<Timeline>;
    terms: ResourceEntry<Terms>;
    privacys: ResourceEntry<Privacy>;
    lineitems: ResourceEntry<LineItem>;
    carousels: ResourceEntry<Carousel>;
    registrations: ResourceEntry<MachineRegistration>;
    warranty: ResourceEntry<WarrantyClaim>;
}

export interface FormField {
    name: string;
    label: string;
    type: string;
    options?: { label: string; value: string | undefined }[];
    placeholder: string;
    defaultValue?: string;
}
