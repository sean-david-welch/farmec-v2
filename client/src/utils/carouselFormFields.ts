import { Carousel } from '../types/miscTypes';

export const getFormFields = (carousel?: Carousel) => [
    {
        name: 'name',
        label: 'Name',
        type: 'text',
        placeholder: 'Enter name',
        defaultValue: carousel?.name,
    },
    {
        name: 'image',
        label: 'Image',
        type: 'file',
        placeholder: 'Upload image',
        defaultValue: carousel?.image,
    },
];
