import styles from '../styles/Machines.module.css';
import utils from '../styles/Utils.module.css';

import { Link } from 'react-router-dom';
import { Machine } from '../types/supplierTypes';

interface Props {
    machines: Machine[];
}

const Machines: React.FC<Props> = props => {
    return (
        <section id="machines">
            <h1 className={utils.sectionHeading}>Machinery</h1>
            {props.machines.map(machine => (
                <div className={styles.machineCard} key={machine.id}>
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
            ))}
        </section>
    );
};

export default Machines;