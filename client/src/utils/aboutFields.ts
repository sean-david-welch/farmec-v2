import { Employee, Timeline, Terms, Privacy } from '../types/aboutTypes';

export const employeeFormFields = (employee?: Employee) => [
    {
        name: 'name',
        label: 'Name',
        type: 'text',
        placeholder: 'Enter name',
        defaultValue: employee?.name,
    },
    {
        name: 'email',
        label: 'Email',
        type: 'email',
        placeholder: 'Enter email',
        defaultValue: employee?.email,
    },
    {
        name: 'role',
        label: 'Role',
        type: 'text',
        placeholder: 'Enter role',
        defaultValue: employee?.role,
    },
    {
        name: 'profile_image',
        label: 'Profile Image',
        type: 'file',
        placeholder: 'Upload profile image',
    },
];

export const timelineFormFields = (timeline?: Timeline) => [
    {
        name: 'title',
        label: 'Title',
        type: 'text',
        placeholder: 'Enter title',
        defaultValue: timeline?.title,
    },
    {
        name: 'date',
        label: 'Date',
        type: 'text',
        placeholder: 'Enter date',
        defaultValue: timeline?.date,
    },
    {
        name: 'body',
        label: 'Body',
        type: 'text',
        placeholder: 'Enter body',
        defaultValue: timeline?.body,
    },
];

export const termsFormFields = (term?: Terms) => [
    {
        name: 'title',
        label: 'Title',
        type: 'text',
        placeholder: 'Enter title',
        defaultValue: term?.title,
    },
    {
        name: 'body',
        label: 'Body',
        type: 'text',
        placeholder: 'Enter body',
        defaultValue: term?.body,
    },
];

export const privacyFormFields = (privacy?: Privacy) => [
    {
        name: 'title',
        label: 'Title',
        type: 'text',
        placeholder: 'Enter title',
        defaultValue: privacy?.title,
    },
    {
        name: 'body',
        label: 'Body',
        type: 'text',
        placeholder: 'Enter body',
        defaultValue: privacy?.body,
    },
];
