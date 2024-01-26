import utils from '../styles/Utils.module.css';

import { useDeleteResource } from '../hooks/genericHooks';

interface ButtonProps {
    id: string;
    route: string;
    queryKey: string;
}

const DeleteButton: React.FC<ButtonProps> = ({ id, route, queryKey }) => {
    const resourceData = { id, route, queryKey };

    const { mutateAsync: deleteResource, isError, error } = useDeleteResource(resourceData);

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
