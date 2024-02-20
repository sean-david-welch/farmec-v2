import utils from '../styles/Utils.module.css';

import { Link } from 'react-router-dom';
import { faArrowRight } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

const NotFound: React.FC = () => {
    const routes = [
        { path: '/', label: 'Home' },
        { path: '/about', label: 'About' },
        { path: '/spareparts', label: 'Spare Parts' },
        { path: '/suppliers', label: 'Suppliers' },
        { path: '/blogs', label: 'Blogs' },
    ];

    return (
        <div className={utils.error}>
            <h1 className={utils.sectionHeading}>Page not found:</h1>
            <p className={utils.subHeading}>Here are some pages you might be interested in:</p>
            <ul className={utils.routesList}>
                {routes.map(route => (
                    <li key={route.path}>
                        <Link to={route.path}>
                            <button className={utils.btn}>
                                {route.label} <FontAwesomeIcon icon={faArrowRight} />
                            </button>
                        </Link>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default NotFound;
