import styles from '../styles/Footer.module.css';
import utils from '../styles/Utils.module.css';

const Footer: React.FC = () => {
    return (
        <footer className={styles.footer}>
            <div className={styles.footerInfo}>
                <a href="/" aria-label="logo button">
                    <img
                        src="https://d3eerclezczw8.cloudfront.net/farmec_images/farmeclogo.webp"
                        alt="Logo"
                        width="250"
                        height="250"
                    />
                </a>

                <ul className={styles.companyInfo}>
                    <li className={styles.infoItem}>Farmec Ireland Ltd.</li>
                    <li className={styles.infoItem}>Clonross, Drumree</li>
                    <li className={styles.infoItem}>Co. Meath, A85 PK30</li>
                    <li className={styles.infoItem}>Tel: 01 - 8259289</li>
                    <li className={styles.infoItem}>Email: info@farmec.ie</li>
                    <li className={styles.infoItem}>
                        <a href={'/about/policy'}>Privacy Policy | Terms of Use</a>
                    </li>
                </ul>
            </div>

            <div className={styles.footerLinks}>
                <div className={styles.accreditation}>
                    <h1 className={utils.mainHeading}>Accreditation</h1>
                    <a href={'https://ftmta.ie/'} target={'_blank'}>
                        <img
                            src="https://d3eerclezczw8.cloudfront.net/farmec_images/ftmta-logo.webp"
                            alt="Logo"
                            width="250"
                            height="250"
                        />
                    </a>
                </div>
                <div className={styles.socialLinks}>
                    <a
                        className={styles.socials}
                        href={'https://www.facebook.com/FarmecIreland/'}
                        target={'_blank'}
                        rel={'noopener noreferrer'}
                        aria-label="Visit our Facebook page">
                        <img src="/icons/facebook.svg" alt="facebook icon" />
                    </a>
                    <a
                        className={styles.socials}
                        href={'https://twitter.com/farmec1?lang=en'}
                        target={'_blank'}
                        rel={'noopener noreferrer'}
                        aria-label="Visit our Twiiter page">
                        <img src="/icons/twitter.svg" alt="facebook icon" />
                    </a>
                </div>
            </div>

            <div className={styles.footerNav}>
                <ul className={styles.navLinks}>
                    <button className={utils.btnFooter}>
                        <a href={'/about'}>Home</a>
                    </button>
                    <button className={utils.btnFooter}>
                        <a href={'/about'}>About</a>
                    </button>
                    <button className={utils.btnFooter}>
                        <a href={'/suppliers'}>Suppliers</a>
                    </button>
                    <button className={utils.btnFooter}>
                        <a href={'/spareparts'}>Spare Parts</a>
                    </button>
                    <button className={utils.btnFooter}>
                        <a href={'/blog'}>Blog</a>
                    </button>
                </ul>
            </div>
        </footer>
    );
};

export default Footer;
