import styles from '../styles/Suppliers.module.css';

import { Component } from 'solid-js';
import { createResource, For } from 'solid-js';

const Suppliers: Component = () => {
    const [suppliers] = createResource(async () => {
        const response = await fetch('http://localhost:8080/api/suppliers');
        return response.json();
    });

    return (
        <>
            <h1>Suppliers</h1>
            <ul>
                <For each={suppliers()}>
                    {supplier => (
                        <div class={styles.supplierCard} id={supplier.name}>
                            <div class={styles.supplierGrid}>
                                <div class={styles.supplierHead}>
                                    <h1 class={styles.mainHeading}>{supplier.name}</h1>
                                    <img
                                        src={supplier.logo_image || '/default.jpg'}
                                        alt={'/default.jpg'}
                                        class={styles.supplierLogo}
                                        width={200}
                                        height={200}
                                    />
                                </div>
                                <img
                                    src={supplier.marketing_image || '/default.jpg'}
                                    alt={'/default.jpg'}
                                    class={styles.supplierImage}
                                    width={550}
                                    height={550}
                                />
                            </div>

                            <div class={styles.supplierInfo}>
                                <p class={styles.supplierDescription}>{supplier.description}</p>
                                <button class={styles.btn}>
                                    <a href={`/suppliers/${supplier.id}`}>Learn More</a>
                                </button>
                            </div>
                        </div>
                    )}
                </For>
            </ul>
        </>
    );
};

export default Suppliers;
