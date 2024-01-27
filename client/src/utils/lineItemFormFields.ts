import { LineItem } from '../types/miscTypes';

export const getFormFields = (lineItem?: LineItem) => [
    {
        name: 'name',
        label: 'Name',
        type: 'text',
        placeholder: 'Enter name',
        defaultValue: lineItem?.name,
    },
    {
        name: 'price',
        label: 'Price',
        type: 'number',
        placeholder: 'Enter price',
        defaultValue: lineItem?.price,
    },
    {
        name: 'image',
        label: 'Image',
        type: 'file',
        placeholder: 'Upload image',
        defaultValue: lineItem?.image,
    },
];
