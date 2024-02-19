import utils from '../../styles/Utils.module.css';
import styles from '../../styles/Suppliers.module.css';

import Error from '../../layouts/Error';
import Videos from '../../templates/Videos';
import Loading from '../../layouts/Loading';
import Machines from '../../templates/Machines';

import { useParams } from 'react-router-dom';
import { Resources } from '../../types/dataTypes';
import { useMultipleResources } from '../../hooks/genericHooks';
import { useSupplierStore, useUserStore } from '../../lib/store';
import { useEffect } from 'react';

const SuppliersDetails: React.FC = () => {
    const { isAdmin } = useUserStore();
    const { suppliers } = useSupplierStore();

    const id = useParams<{ id: string }>().id as string;

    const resourceKeys: (keyof Resources)[] = ['suppliers', 'supplierMachine', 'videos'];
    const { data, isLoading, isError } = useMultipleResources(id, resourceKeys);

    if (isError) return <Error />;
    if (isLoading) return <Loading />;

    useEffect(() => {}, [id]);

    const [supplier, machines, videos] = data;

    return (
        <section id="supplierDetail">
            {supplier && (
                <>
                    <div className={styles.supplierHeading}>
                        <h1 className={utils.sectionHeading}>{supplier.name}</h1>
                    </div>

                    {machines.length > 0 && (
                        <div className={utils.index}>
                            <h1 className={utils.indexHeading}>Suppliers</h1>
                            {machines.map((link: { name: string }) => (
                                <a key={link.name} href={`#${link.name}`}>
                                    <h1 className={utils.indexItem}>{link.name}</h1>
                                </a>
                            ))}
                        </div>
                    )}

                    <div className={styles.supplierDetail}>
                        <img
                            src={supplier.marketing_image ?? '/default.jpg'}
                            alt={'/dafault.jpg'}
                            className={styles.supplierImage}
                            width={750}
                            height={750}
                        />

                        <p className={styles.supplierDescription}>{supplier.description}</p>
                    </div>
                </>
            )}

            {machines.length > 0 && <Machines machines={machines} isAdmin={isAdmin} />}
            {videos.length > 0 && <Videos suppliers={suppliers} videos={videos} isAdmin={isAdmin} />}
        </section>
    );
};

export default SuppliersDetails;
