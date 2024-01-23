import { Component, For, createResource } from 'solid-js';
import AccountButton from '../components/AccountButton';

import styles from '../styles/Header.module.css';
import type Supplier from '../types/supplier';
import { useLocation } from '@solidjs/router';

const Header: Component = () => {
    const location = useLocation();
    const isHomepage = () => location.pathname === '/';

    const [suppliers] = createResource<Supplier[]>(async () => {
        return fetch(`http://localhost:8080/api/suppliers`).then(response => response.json());
    });

    return (
        <nav class={isHomepage() ? styles.transparentNav : styles.navbar}>
            <a href="/" aria-label="logo button">
                <img
                    src="https://farmec-bucket.s3.eu-west-1.amazonaws.com/farmec_images/farmeclogo.webp"
                    alt="Logo"
                    width="250"
                    height="250"
                />
            </a>

            <ul class={styles.navList}>
                <li class={styles.navItem}>
                    <a href="/about" style={styles.navListItem}>
                        About Us
                    </a>
                    <ul class={styles.navDrop}>
                        <li class={styles.navDropItem}>
                            <a href="/about">Staff & Management</a>
                        </li>
                        <li class={styles.navDropItem}>
                            <a href="/about#timeline">Company History</a>
                        </li>
                        <li class={styles.navDropItem}>
                            <a href="/about/policies">Terms of Use</a>
                        </li>
                    </ul>
                </li>

                <li class={styles.navItem}>
                    <a href="/suppliers" style={styles.navListItem}>
                        Suppliers
                    </a>
                    <ul class={styles.navDrop}>
                        <For each={suppliers()}>
                            {supplier => (
                                <li class={styles.navDropItem}>
                                    <a href={`/suppliers/${supplier.id}`}>{supplier.name}</a>
                                </li>
                            )}
                        </For>
                    </ul>
                </li>

                <li class={styles.navItem}>
                    <a href="/spareparts" style={styles.navListItem}>
                        Spareparts
                    </a>
                    <ul class={styles.navDrop}>
                        <For each={suppliers()}>
                            {supplier => (
                                <li class={styles.navDropItem}>
                                    <a href={`/spareparts/${supplier.id}`}>{supplier.name}</a>
                                </li>
                            )}
                        </For>
                    </ul>
                </li>

                <li class={styles.navItem}>
                    <a href="/blogs" style={styles.navListItem}>
                        Blog
                    </a>
                    <ul class={styles.navDrop}>
                        <li class={styles.navDropItem}>
                            <a href="/blog">Latest Posts</a>
                        </li>
                        <li class={styles.navDropItem}>
                            <a href="/blog/exhibitions">Exhibition Information</a>
                        </li>
                    </ul>
                </li>

                <li class={styles.navItem}>
                    <a href="/#contact" style={styles.navListItem}>
                        Contact
                    </a>
                </li>

                <AccountButton />
            </ul>
        </nav>
    );
};

export default Header;
