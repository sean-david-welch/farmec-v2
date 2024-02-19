import utils from '../styles/Utils.module.css';
import styles from '../styles/Footer.module.css';

import ToTopButton from '../components/ToTopButton';

import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTwitter } from '@fortawesome/free-brands-svg-icons/faTwitter';
import { faFacebook } from '@fortawesome/free-brands-svg-icons/faFacebook';
import { faRightToBracket } from '@fortawesome/free-solid-svg-icons/faRightToBracket';

const Footer: React.FC = () => {
    return (
        <footer className={styles.footer}>
            <div className={styles.footerInfo}>
                <Link to="/" aria-label="logo button">
                    <img
                        src="https://www.farmec.ie/farmec_images/farmeclogo.webp"
                        alt="Logo"
                        width="180"
                        height="180"
                    />
                </Link>

                <ul className={styles.companyInfo}>
                    <li className={styles.infoItem}>Farmec Ireland Ltd.</li>
                    <li className={styles.infoItem}>Clonross, Drumree</li>
                    <li className={styles.infoItem}>Co. Meath, A85 PK30</li>
                    <li className={styles.infoItem}>Tel: 01 - 8259289</li>
                    <li className={styles.infoItem}>Email: info@farmec.ie</li>
                    <li className={styles.infoItem}>
                        <Link to={'/about/policies'}>Privacy Policy | Terms of Use</Link>
                    </li>
                </ul>
            </div>

            <div className={styles.footerLinks}>
                <h1 className={utils.mainHeading}>Accreditation</h1>
                <Link to={'https://ftmta.ie/'} target={'_blank'}>
                    <img
                        src="https://www.farmec.ie/farmec_images/ftmta-logo.webp"
                        alt="Logo"
                        width="180"
                        height="180"
                    />
                </Link>

                <div className={styles.socialLinks}>
                    <Link
                        className={styles.socials}
                        to={'https://www.facebook.com/FarmecIreland/'}
                        target={'_blank'}
                        rel={'noopener noreferrer'}
                        aria-label="Visit our Facebook page">
                        <FontAwesomeIcon icon={faFacebook} />
                    </Link>
                    <Link
                        className={styles.socials}
                        to={'https://twitter.com/farmec1?lang=en'}
                        target={'_blank'}
                        rel={'noopener noreferrer'}
                        aria-label="Visit our Twiiter page">
                        <FontAwesomeIcon icon={faTwitter} />
                    </Link>
                </div>
            </div>

            <div className={styles.footerNav}>
                <ul className={styles.navLinks}>
                    <button className={utils.btnFooter}>
                        <Link to={'/about'}>
                            Home <FontAwesomeIcon icon={faRightToBracket} />
                        </Link>
                    </button>
                    <button className={utils.btnFooter}>
                        <Link to={'/about'}>
                            About <FontAwesomeIcon icon={faRightToBracket} />
                        </Link>
                    </button>
                    <button className={utils.btnFooter}>
                        <Link to={'/suppliers'}>
                            Suppliers <FontAwesomeIcon icon={faRightToBracket} />
                        </Link>
                    </button>
                    <button className={utils.btnFooter}>
                        <Link to={'/spareparts'}>
                            Spare Parts <FontAwesomeIcon icon={faRightToBracket} />
                        </Link>
                    </button>
                    <button className={utils.btnFooter}>
                        <Link to={'/blog'}>
                            Blog <FontAwesomeIcon icon={faRightToBracket} />
                        </Link>
                    </button>
                </ul>
            </div>

            <ToTopButton />
        </footer>
    );
};

export default Footer;
