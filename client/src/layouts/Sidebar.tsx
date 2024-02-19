import styles from '../styles/Sidebar.module.css';

import MobileButton from '../components/MobileLogin';

import { Link } from 'react-router-dom';
import { useState, Fragment } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
    faBars,
    faHome,
    faX,
    faIndustry,
    faCircleInfo,
    faCog,
    faBlog,
    faQuestionCircle,
} from '@fortawesome/free-solid-svg-icons';

const Sidebar = () => {
    const isHomepage = () => location.pathname === '/';

    const [isOpen, setIsOpen] = useState<boolean>(false);

    const renderContent = () => {
        if (!isOpen) {
            const navIconContent = (
                <div className={styles.navIcon} onClick={() => setIsOpen(!isOpen)}>
                    <Link to="/" aria-label="logo button">
                        <img
                            src="https://www.farmec.ie/farmec_images/farmeclogo.webp"
                            alt="Logo"
                            width="200"
                            height="200"
                        />
                    </Link>
                    <FontAwesomeIcon icon={faBars} className={styles.navigation} />
                </div>
            );

            return isHomepage() ? navIconContent : <div className={styles.backdrop}>{navIconContent}</div>;
        } else {
            return (
                <nav className={isOpen ? styles.sideNavOpen : styles.sideNavClosed}>
                    <div className={styles.navIcon} onClick={() => setIsOpen(!isOpen)}>
                        <Link to="/" aria-label="logo button">
                            <img
                                src="https://www.farmec.ie/farmec_images/farmeclogo.webp"
                                alt="Logo"
                                width="200"
                                height="200"
                            />
                        </Link>
                        <FontAwesomeIcon icon={faX} className={styles.navigation} />
                    </div>
                    <ul className={styles.navList}>
                        <Link className={styles.navItem} to={'/'} onClick={() => setIsOpen(false)}>
                            <FontAwesomeIcon icon={faHome} />
                            Home
                        </Link>
                        <Link className={styles.navItem} to={'/about'} onClick={() => setIsOpen(false)}>
                            <FontAwesomeIcon icon={faCircleInfo} />
                            About
                        </Link>
                        <Link className={styles.navItem} to={'/about/policies'} onClick={() => setIsOpen(false)}>
                            <FontAwesomeIcon icon={faQuestionCircle} />
                            Terms of use
                        </Link>
                        <Link className={styles.navItem} to={'/suppliers'} onClick={() => setIsOpen(false)}>
                            <FontAwesomeIcon icon={faIndustry} />
                            Suppliers
                        </Link>
                        <Link className={styles.navItem} to={'/spareparts'} onClick={() => setIsOpen(false)}>
                            <FontAwesomeIcon icon={faCog} />
                            Spareparts
                        </Link>
                        <Link className={styles.navItem} to={'/blogs'} onClick={() => setIsOpen(false)}>
                            <FontAwesomeIcon icon={faBlog} />
                            Blog
                        </Link>
                        <Link to={'/blog/exhibitions'} className={styles.navItem} onClick={() => setIsOpen(false)}>
                            <FontAwesomeIcon icon={faQuestionCircle} />
                            Exhibitions
                        </Link>
                        <MobileButton onClick={() => setIsOpen(false)} />
                    </ul>
                </nav>
            );
        }
    };

    return <Fragment>{renderContent()}</Fragment>;
};

export default Sidebar;
