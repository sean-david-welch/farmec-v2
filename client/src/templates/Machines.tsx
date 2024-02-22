import styles from '../styles/Machines.module.css';
import utils from '../styles/Utils.module.css';

import { Link } from 'react-router-dom';
import { Machine } from '../types/supplierTypes';
import { Fragment } from 'react';

import Loading from '../layouts/Loading';
import MachineFrom from '../forms/MachineForm';
import DeleteButton from '../components/DeleteButton';

import { useSupplierStore } from '../lib/store';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faRightToBracket } from '@fortawesome/free-solid-svg-icons/faRightToBracket';

interface Props {
    machines: Machine[];
    isAdmin: boolean;
}

const Machines: React.FC<Props> = ({ machines, isAdmin }) => {
    const { suppliers } = useSupplierStore();

    if (!suppliers) {
        return <Loading />;
    }

    return (
        <section id="machines">
            <h1 className={utils.sectionHeading}>Machinery</h1>
            {machines.map(machine => (
                <Fragment key={machine.id}>
                    <div className={styles.machineCard} id={machine.name}>
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
                                        <FontAwesomeIcon icon={faRightToBracket} />
                                    </Link>
                                </button>
                            </div>
                        </div>
                    </div>

                    {isAdmin && machine.id && (
                        <div className={utils.optionsBtn}>
                            <MachineFrom id={machine.id} machine={machine} suppliers={suppliers} />
                            <DeleteButton id={machine.id} resourceKey="machines" />
                        </div>
                    )}
                </Fragment>
            ))}

            {isAdmin && <MachineFrom suppliers={suppliers} />}
        </section>
    );
};

export default Machines;
