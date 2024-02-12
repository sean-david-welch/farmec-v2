import utils from '../../styles/Utils.module.css';

import Products from '../../templates/Products';

import { useParams } from 'react-router-dom';
import { Resources } from '../../types/dataTypes';
import { useUserStore } from '../../lib/store';
import { useMultipleResources } from '../../hooks/genericHooks';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faRightToBracket } from '@fortawesome/free-solid-svg-icons/faRightToBracket';

const MachineDetail: React.FC = () => {
    const { isAdmin } = useUserStore();

    const id = useParams<{ id: string }>().id as string;

    const resourceKeys: (keyof Resources)[] = ['machines', 'products'];
    const { data, isLoading } = useMultipleResources(id, resourceKeys);

    if (!id) {
        return <div>Error: No supplier ID provided</div>;
    }

    if (isError) return <Error />;
    if (isLoading) return <Loading />;

    const [machine, products] = data;

    return (
        <section id="machineDetai">
            <h1 className={utils.sectionHeading}>Products</h1>
            {products && (
                <div className={utils.index}>
                    <h1 className={utils.indexHeading}>Suppliers</h1>
                    {products.map((link: { name: string }) => (
                        <a key={link.name} href={`#${link.name}`}>
                            <h1 className={utils.indexItem}>{link.name}</h1>
                        </a>
                    ))}
                    <button className={utils.btn}>
                        <a href={machine.machine_link || '#'} target="_blank">
                            Supplier Website
                            <FontAwesomeIcon icon={faRightToBracket} />
                        </a>
                    </button>
                </div>
            )}
            {products && <Products id={id} products={products} isAdmin={isAdmin} />}
        </section>
    );
};

export default MachineDetail;
