export default interface Blog {
    id?: string;
    title: string;
    date: string;
    main_image: string;
    subheading: string;
    body: string;
    created?: string;
}

export interface Exhibition {
    id?: string;
    title: string;
    date: string;
    location: string;
    info: string;
    created?: string;
}
