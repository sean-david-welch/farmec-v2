import utils from '../../styles/Machines.module.css';

import { useParams } from 'react-router-dom';
import { useMachineDetails } from '../../hooks/supplierHooks';
import Products from '../../templates/Products';

const MachineDetail = async () => {
    const params = useParams<{ id: string }>();

    if (!params.id) {
        return <div>Error: No supplier ID provided</div>;
    }

    const { machine, products } = useMachineDetails(params.id);

    if (!machine) {
        return <div>loading...</div>;
    }

    return (
        <section id="machineDetai">
            <h1 className={utils.sectionHeading}>Products</h1>
            {products?.data && (
                <div className={utils.index}>
                    <h1 className={utils.indexHeading}>Suppliers</h1>
                    {products.data?.map(link => (
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
