import utils from '../styles/Utils.module.css';

import { useDeleteResource } from '../hooks/genericHooks';
import { Resources } from '../types/dataTypes';

interface ButtonProps {
    id: string;
    resourceKey: keyof Resources;
}

const DeleteButton: React.FC<ButtonProps> = ({ id, resourceKey }) => {
    const { mutateAsync: deleteResource, isError, error } = useDeleteResource(resourceKey, id);

    async function handleSubmit(event: React.MouseEvent<HTMLButtonElement>) {
        event.preventDefault();

        try {
            await deleteResource();
        } catch (error) {
            console.error('Error creating supplier:', error);
        }
    }

    return (
        <button className={utils.btnForm} onClick={handleSubmit}>
            <img src="/icons/trash.svg" alt="trash icon" />
            {isError && <p>Error: {error.message}</p>}
        </button>
    );
};

export default DeleteButton;
