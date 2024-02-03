import { Blog, Exhibition } from '../types/blogTypes';

export const blogFormFields = (blog?: Blog) => [
    {
        name: 'title',
        label: 'Title',
        type: 'text',
        placeholder: 'Enter title',
        defaultValue: blog?.title,
    },
    {
        name: 'date',
        label: 'Date',
        type: 'text',
        placeholder: 'Enter date',
        defaultValue: blog?.date,
    },
    {
        name: 'main_image',
        label: 'Main Image',
        type: 'file',
        placeholder: 'Upload main image',
        defaultValue: blog?.main_image,
    },
    {
        name: 'subheading',
        label: 'Subheading',
        type: 'text',
        placeholder: 'Enter subheading',
        defaultValue: blog?.subheading,
    },
    {
        name: 'body',
        label: 'Body',
        type: 'text',
        placeholder: 'Enter body',
        defaultValue: blog?.body,
    },
];

export const exhibitionFormFields = (exhibiton?: Exhibition) => [
    {
        name: 'title',
        label: 'Title',
        type: 'text',
        placeholder: 'Enter title',
        defaultValue: exhibiton?.title,
    },
    {
        name: 'date',
        label: 'Date',
        type: 'text',
        placeholder: 'Enter date',
        defaultValue: exhibiton?.title,
    },
    {
        name: 'location',
        label: 'Location',
        type: 'text',
        placeholder: 'Enter location',
        defaultValue: exhibiton?.title,
    },
    {
        name: 'info',
        label: 'Info',
        type: 'text',
        placeholder: 'Enter info',
        defaultValue: exhibiton?.title,
    },
];
