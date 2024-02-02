import utils from '../../styles/Machines.module.css';

import { useParams } from 'react-router-dom';

import Products from '../../templates/Products';
import { Resources } from '../../types/dataTypes';
import { useMultipleResources } from '../../hooks/genericHooks';

const MachineDetail: React.FC = () => {
    const params = useParams<{ id: string }>();

    if (!params.id) {
        return <div>Error: No supplier ID provided</div>;
    }

    const resourceKeys: (keyof Resources)[] = ['machines', 'products'];
    const { data, isLoading } = useMultipleResources(params.id, resourceKeys);

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
                    {products.data?.map((link: { name: string }) => (
                        <a key={link.name} href={`#${link.name}`}>
                            <h1 className="indexItem">{link.name}</h1>
                        </a>
                    ))}
                    <button className={utils.btn}>
                        <a href={machine.data?.machine_link || '#'} target="_blank">
                            Supplier Website
                            <img src="/icons/right-bracket.svg" alt="bracket-right" />
                        </a>
                    </button>
                </div>
            )}
            {products?.data && <Products products={products?.data} />}
        </section>
    );
};

export default MachineDetail;
