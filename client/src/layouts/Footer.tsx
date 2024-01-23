import styles from '../styles/Footer.module.css';
import utils from '../styles/Utils.module.css';

import { Component } from 'solid-js';

const Footer: Component = () => {
    return (
        <footer class={styles.footer}>
            <div class={styles.footerInfo}>
                <a href="/" aria-label="logo button">
                    <img
                        src="https://res.cloudinary.com/dgpquyhuy/image/upload/v1691492819/farmeclogo.png"
                        alt="Logo"
                        width="250"
                        height="250"
                    />
                </a>

                <ul class={styles.companyInfo}>
                    <li class={styles.infoItem}>Farmec Ireland Ltd.</li>
                    <li class={styles.infoItem}>Clonross, Drumree</li>
                    <li class={styles.infoItem}>Co. Meath, A85 PK30</li>
                    <li class={styles.infoItem}>Tel: 01 - 8259289</li>
                    <li class={styles.infoItem}>Email: info@farmec.ie</li>
                    <li class={styles.infoItem}>
                        <a href={'/about/policy'}>Privacy Policy | Terms of Use</a>
                    </li>
                </ul>
            </div>

            <div class={styles.footerLinks}>
                <div class={styles.accreditation}>
                    <h1 class={utils.mainHeading}>Accreditation</h1>
                    <a href={'https://ftmta.ie/'} target={'_blank'}>
                        <img
                            src="https://res.cloudinary.com/dgpquyhuy/image/upload/v1691492819/ftmta-logo.png"
                            alt="Logo"
                            width="250"
                            height="250"
                        />
                    </a>
                </div>
                <div class={styles.socialLinks}>
                    <a
                        class={styles.socials}
                        href={'https://www.facebook.com/FarmecIreland/'}
                        target={'_blank'}
                        rel={'noopener noreferrer'}
                        aria-label="Visit our Facebook page">
                        <img src="/icons/facebook.svg" alt="facebook icon" />
                    </a>
                    <a
                        class={styles.socials}
                        href={'https://twitter.com/farmec1?lang=en'}
                        target={'_blank'}
                        rel={'noopener noreferrer'}
                        aria-label="Visit our Twiiter page">
                        <img src="/icons/twitter.svg" alt="facebook icon" />
                    </a>
                </div>
            </div>

            <div class={styles.footerNav}>
                <ul class={styles.navLinks}>
                    <button class={utils.btnFooter}>
                        <a href={'/about'}>Home</a>
                    </button>
                    <button class={utils.btnFooter}>
                        <a href={'/about'}>About</a>
                    </button>
                    <button class={utils.btnFooter}>
                        <a href={'/suppliers'}>Suppliers</a>
                    </button>
                    <button class={utils.btnFooter}>
                        <a href={'/spareparts'}>Spare Parts</a>
                    </button>
                    <button class={utils.btnFooter}>
                        <a href={'/blog'}>Blog</a>
                    </button>
                </ul>
            </div>
        </footer>
    );
};

export default Footer;
