import styles from '../styles/Footer.module.css';
import utils from '../styles/Utils.module.css';

import { Link } from 'react-router-dom';

const Footer: React.FC = () => {
    return (
        <footer className={styles.footer}>
            <div className={styles.footerInfo}>
                <Link to="/" aria-label="logo button">
                    <img
                        src="https://d3eerclezczw8.cloudfront.net/farmec_images/farmeclogo.webp"
                        alt="Logo"
                        width="250"
                        height="250"
                    />
                </Link>

                <ul className={styles.companyInfo}>
                    <li className={styles.infoItem}>Farmec Ireland Ltd.</li>
                    <li className={styles.infoItem}>Clonross, Drumree</li>
                    <li className={styles.infoItem}>Co. Meath, A85 PK30</li>
                    <li className={styles.infoItem}>Tel: 01 - 8259289</li>
                    <li className={styles.infoItem}>Email: info@farmec.ie</li>
                    <li className={styles.infoItem}>
                        <Link to={'/about/policy'}>Privacy Policy | Terms of Use</Link>
                    </li>
                </ul>
            </div>

            <div className={styles.footerLinks}>
                <div className={styles.accreditation}>
                    <h1 className={utils.mainHeading}>Accreditation</h1>
                    <Link to={'https://ftmta.ie/'} target={'_blank'}>
                        <img
                            src="https://d3eerclezczw8.cloudfront.net/farmec_images/ftmta-logo.webp"
                            alt="Logo"
                            width="250"
                            height="250"
                        />
                    </Link>
                </div>
                <div className={styles.socialLinks}>
                    <Link
                        className={styles.socials}
                        to={'https://www.facebook.com/FarmecIreland/'}
                        target={'_blank'}
                        rel={'noopener noreferrer'}
                        aria-label="Visit our Facebook page">
                        <img src="/icons/facebook.svg" alt="facebook icon" />
                    </Link>
                    <Link
                        className={styles.socials}
                        to={'https://twitter.com/farmec1?lang=en'}
                        target={'_blank'}
                        rel={'noopener noreferrer'}
                        aria-label="Visit our Twiiter page">
                        <img src="/icons/twitter.svg" alt="facebook icon" />
                    </Link>
                </div>
            </div>

            <div className={styles.footerNav}>
                <ul className={styles.navLinks}>
                    <button className={utils.btnFooter}>
                        <Link to={'/about'}>Home</Link>
                    </button>
                    <button className={utils.btnFooter}>
                        <Link to={'/about'}>About</Link>
                    </button>
                    <button className={utils.btnFooter}>
                        <Link to={'/suppliers'}>Suppliers</Link>
                    </button>
                    <button className={utils.btnFooter}>
                        <Link to={'/spareparts'}>Spare Parts</Link>
                    </button>
                    <button className={utils.btnFooter}>
                        <Link to={'/blog'}>Blog</Link>
                    </button>
                </ul>
            </div>
        </footer>
    );
};

export default Footer;
