import styles from '../styles/Machines.module.css';
import utils from '../styles/Utils.module.css';

import Machine from '../types/machine';

import { Component, For } from 'solid-js';

interface Props {
    machines: Machine[];
}

const Machines: Component<Props> = props => {
    return (
        <section id="machines">
            <h1 class={utils.sectionHeading}>Machinery</h1>
            <For each={props.machines}>
                {machine => (
                    <div class={styles.machineCard} id={machine.name || ''}>
                        <div class={styles.machineGrid}>
                            <img
                                src={machine.machine_image || '/default.jpg'}
                                alt={machine.name || 'Default Image'}
                                class={styles.machineImage}
                                width={600}
                                height={600}
                            />
                            <div class={styles.machineInfo}>
                                <h1 class={utils.mainHeading}>{machine.name}</h1>
                                <p class={utils.paragraph}>{machine.description}</p>
                                <button class={utils.btn}>
                                    <a href={`/machines/${machine.id}`}>
                                        View Products
                                        <i aria-label="icon">
                                            <img src="/icons/right-bracket.svg" alt="bracket-right" />
                                        </i>
                                    </a>
                                </button>
                            </div>
                        </div>
                    </div>
                )}
            </For>
        </section>
    );
};

export default Machines;
