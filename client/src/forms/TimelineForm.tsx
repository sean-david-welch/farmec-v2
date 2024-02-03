import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';

import { useState } from 'react';
import { Timeline } from '../types/aboutTypes';

import { useMutateResource } from '../hooks/genericHooks';
import { timelineFormFields } from '../utils/aboutFields';

interface Props {
    id?: string;
    timeline?: Timeline;
}

const TimelineForm: React.FC<Props> = ({ id, timeline }) => {
    const [showForm, setShowForm] = useState(false);
    const formFields = timeline ? timelineFormFields(timeline) : timelineFormFields();

    const {
        mutateAsync: createTimeline,
        isError: isCreateError,
        error: createError,
    } = useMutateResource<Timeline>('timelines');

    const {
        mutateAsync: updateTimeline,
        isError: isUpdateError,
        error: updateError,
    } = useMutateResource<Timeline>('timelines', id);

    const error = id ? updateError : createError;
    const isError = id ? isUpdateError : isCreateError;
    const submitTimeline = id ? updateTimeline : createTimeline;

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        const formData = new FormData(event.currentTarget as HTMLFormElement);

        const body: Timeline = {
            title: formData.get('title') as string,
            date: formData.get('date') as string,
            body: formData.get('body') as string,
        };

        try {
            const response = await submitTimeline(body);
            response.ok ? setShowForm(false) : console.error('failed with response:', response);
        } catch (error) {
            console.error('error creating timeline', error);
        }
    }

    return (
        <section id="form">
            <button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
                {id ? <img src="/icons/edit.svg" alt="edit button" /> : 'Create Timeline'}
            </button>

            <FormDialog visible={showForm} onClose={() => setShowForm(false)}>
                <form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
                    <h1 className={utils.mainHeading}>Supplier Form</h1>
                    {formFields.map((field) => (
                        <div key={field.name}>
                            <label htmlFor={field.name}>{field.label}</label>
                            <input
                                type={field.type}
                                name={field.name}
                                id={field.name}
                                placeholder={field.placeholder}
                                defaultValue={field.defaultValue}
                            />
                        </div>
                    ))}
                    <button className={utils.btnForm} type="submit">
                        Submit
                    </button>
                </form>

                {isError && <p>Error: {error?.message}</p>}
            </FormDialog>
        </section>
    );
};

export default TimelineForm;
