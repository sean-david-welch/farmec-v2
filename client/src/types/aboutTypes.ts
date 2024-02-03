export interface Employee {
    id?: string;
    name: string;
    email: string;
    role: string;
    profile_image: string;
    created?: string;
}

export interface Timeline {
    id?: string;
    title: string;
    date: string;
    body: string;
    created?: string;
}

export interface Privacy {
    id?: string;
    title: string;
    body: string;
    created?: string;
}

export interface Terms {
    id?: string;
    title: string;
    body: string;
    created?: string;
}
