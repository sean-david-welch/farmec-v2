import utils from '../../styles/Machines.module.css';

import Products from '../../templates/Products';

import { useParams } from 'react-router-dom';
import { Resources } from '../../types/dataTypes';
import { useUserStore } from '../../lib/store';
import { useMultipleResources } from '../../hooks/genericHooks';

const MachineDetail: React.FC = () => {
    const { isAdmin } = useUserStore();

    const id = useParams<{ id: string }>().id as string;

    const resourceKeys: (keyof Resources)[] = ['machines', 'products'];
    const { data, isLoading } = useMultipleResources(id, resourceKeys);

    if (!id) {
        return <div>Error: No supplier ID provided</div>;
    }

    if (isLoading) {
        return <div>loading...</div>;
    }

    const [machine, products] = data;

    return (
        <section id="machineDetai">
            <h1 className={utils.sectionHeading}>Products</h1>
            {products && (
                <div className={utils.index}>
                    <h1 className={utils.indexHeading}>Suppliers</h1>
                    {products.map((link: { name: string }) => (
                        <a key={link.name} href={`#${link.name}`}>
                            <h1 className="indexItem">{link.name}</h1>
                        </a>
                    ))}
                    <button className={utils.btn}>
                        <a href={machine.machine_link || '#'} target="_blank">
                            Supplier Website
                            <img src="/icons/right-bracket.svg" alt="bracket-right" />
                        </a>
                    </button>
                </div>
            )}
            {products && <Products products={products} isAdmin={isAdmin} />}
        </section>
    );
};

export default MachineDetail;
