import utils from '../styles/Utils.module.css';

import { useDeleteResource } from '../hooks/genericHooks';
import { Resources } from '../types/dataTypes';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash } from '@fortawesome/free-solid-svg-icons/faTrash';
import { useNavigate } from 'react-router-dom';
import Loading from '../layouts/Loading';

interface ButtonProps {
    id: string;
    resourceKey: keyof Resources;
    navigateBack?: boolean;
}

const DeleteButton: React.FC<ButtonProps> = ({ id, resourceKey, navigateBack }) => {
    const navigate = useNavigate();
    const { mutateAsync: deleteResource, isError, error, isPending } = useDeleteResource(resourceKey, id);

    if (isPending) return <Loading />;

    async function handleSubmit(event: React.MouseEvent<HTMLButtonElement>) {
        event.preventDefault();

        try {
            if (navigateBack) navigate('/');
            await deleteResource();
        } catch (error) {
            console.error('Error deleting resource:', error);
        }
    }

    return (
        <button className={utils.btnForm} onClick={handleSubmit}>
            <FontAwesomeIcon icon={faTrash} />
            {isError && <p>Error: {error.message}</p>}
        </button>
    );
};

export default DeleteButton;
