import utils from '../styles/Utils.module.css';
import styles from '../styles/Suppliers.module.css';

import { Link } from 'react-router-dom';
import { Supplier } from '../types/supplierTypes';
import { useUserStore } from '../lib/store';
import { useGetResource } from '../hooks/genericHooks';

const SpareParts: React.FC = () => {
    const { isAdmin } = useUserStore();

    const suppliers = useGetResource<Supplier[]>('suppliers');

    return (
        <section id="SpareParts">
            <h1 className={utils.sectionHeading}>Spare-Parts</h1>

            {suppliers.data && (
                <div className={utils.index}>
                    <h1 className={utils.indexHeading}>Suppliers</h1>
                    {suppliers.data.map(link => (
                        <Link key={link.name} to={`#${link.name}`}>
                            <h1 className="indexItem">{link.name}</h1>
                        </Link>
                    ))}
                </div>
            )}
            {suppliers.data?.map(supplier => (
                <div className={styles.supplierCard} key={supplier.id}>
                    <h1 className={utils.mainHeading}>{supplier.name}</h1>
                    <img src={supplier.logo_image ?? '/default.jpg'} alt="Supplier logo" width={200} height={200} />
                    <button className={utils.btn}>
                        <Link to={`/spareparts/${supplier.id}`}>
                            Spare-Parts
                            <img src="/icons/right-bracket.svg" alt="bracket-right" />
                        </Link>
                    </button>
                </div>
            ))}
            {/* {isAdmin && <SparePartsForm />} */}
        </section>
    );
};

export default SpareParts;
