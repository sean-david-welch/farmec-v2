import styles from '../styles/Suppliers.module.css';
import SupplierForm from '../forms/SupplierForm';

import { useSuppliers } from '../hooks/supplierHooks';
import { Link } from 'react-router-dom';
import { useUserStore } from '../utils/context';
import DeleteButton from '../components/DeleteButton';

const Suppliers: React.FC = () => {
    const { isAdmin } = useUserStore();
    const suppliers = useSuppliers();

    if (suppliers.isLoading) {
        return <div>Loading...</div>;
    }

    if (suppliers.isError) {
        return <div>Error loading data</div>;
    }

    return (
        <section id="suppliers">
            <h1>Suppliers</h1>
            {suppliers.data && (
                <>
                    {suppliers.data.map(supplier => (
                        <>
                            <div className={styles.supplierCard} id={supplier.name} key={supplier.id}>
                                <div className={styles.supplierGrid}>
                                    <div className={styles.supplierHead}>
                                        <h1 className={styles.mainHeading}>{supplier.name}</h1>
                                        <img
                                            src={supplier.logo_image || '/default.jpg'}
                                            alt={'/default.jpg'}
                                            className={styles.supplierLogo}
                                            width={200}
                                            height={200}
                                        />
                                    </div>
                                    <img
                                        src={supplier.marketing_image || '/default.jpg'}
                                        alt={'/default.jpg'}
                                        className={styles.supplierImage}
                                        width={550}
                                        height={550}
                                    />
                                </div>
                                <div className={styles.supplierInfo}>
                                    <p className={styles.supplierDescription}>{supplier.description}</p>
                                    <button className={styles.btn}>
                                        <Link to={`/suppliers/${supplier.id}`}>Learn More</Link>
                                    </button>
                                </div>
                            </div>

                            {isAdmin && supplier.id && (
                                <div className="">
                                    <SupplierForm id={supplier.id} />
                                    <DeleteButton id={supplier.id} route="suppliers" queryKey="suppliers" />
                                </div>
                            )}
                        </>
                    ))}
                </>
            )}

            {isAdmin && <SupplierForm />}
        </section>
    );
};

export default Suppliers;
