import styles from '../styles/Machines.module.css';
import utils from '../styles/Utils.module.css';

import { Link } from 'react-router-dom';
import { Machine, Supplier } from '../types/supplierTypes';
import { Fragment } from 'react';
import { useUserStore } from '../lib/store';
import MachineFrom from '../forms/MachineForm';
import DeleteButton from '../components/DeleteButton';
import { useGetResource } from '../hooks/genericHooks';

interface Props {
    machines: Machine[];
}

const Machines: React.FC<Props> = ({ machines }) => {
    const isAdmin = useUserStore();

    const suppliers = useGetResource<Supplier[]>('suppliers');

    if (!suppliers.data) {
        return <div>Loading...</div>;
    }

    return (
        <section id="machines">
            <h1 className={utils.sectionHeading}>Machinery</h1>
            {machines.map((machine) => (
                <Fragment key={machine.id}>
                    <div className={styles.machineCard}>
                        <div className={styles.machineGrid}>
                            <img
                                src={machine.machine_image || '/default.jpg'}
                                alt={machine.name || 'Default Image'}
                                className={styles.machineImage}
                                width={600}
                                height={600}
                            />
                            <div className={styles.machineInfo}>
                                <h1 className={utils.mainHeading}>{machine.name}</h1>
                                <p className={utils.paragraph}>{machine.description}</p>
                                <button className={utils.btn}>
                                    <Link to={`/machines/${machine.id}`}>
                                        View Products
                                        <i aria-label="icon">
                                            <img src="/icons/right-bracket.svg" alt="bracket-right" />
                                        </i>
                                    </Link>
                                </button>
                            </div>
                        </div>
                    </div>

                    {isAdmin && machine.id && (
                        <div className={utils.optionsBtn}>
                            <MachineFrom id={machine.id} machine={machine} suppliers={suppliers.data} />
                            <DeleteButton id={machine.id} resourceKey="machines" />
                        </div>
                    )}
                </Fragment>
            ))}

            {isAdmin && <MachineFrom suppliers={suppliers.data} />}
        </section>
    );
};

export default Machines;
