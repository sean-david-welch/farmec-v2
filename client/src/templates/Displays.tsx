import styles from '../styles/Info.module.css';
import utils from '../styles/Utils.module.css';

import { Link } from 'react-router-dom';
import { statsItems, specialsItems } from '../components/displaysInfo';

const Displays: React.FC = () => {
    return (
        <section id="Info">
            <div className={styles.infoSection}>
                <h1 className={utils.sectionHeading}>Farmec At A Glance:</h1>
                <p className={utils.subHeading}>
                    This is a Quick Look at what Separates us from our Competitors
                </p>
                <div className={styles.stats}>
                    {statsItems.map((item, index) => (
                        <Link to={item.link} key={index}>
                            <ul className={styles.statList}>
                                <li className={styles.statListItem}>{item.title}</li>
                                <li className={styles.statListItem}>{item.icon}</li>
                                <li className={styles.statListItem}>{item.description}</li>
                            </ul>
                        </Link>
                    ))}
                </div>
            </div>
            <div className={styles.infoSection}>
                <h1 className={utils.sectionHeading}>Farmec At A Glance:</h1>
                <p className={utils.subHeading}>
                    This is a Quick Look at what Separates us from our Competitors
                </p>
                <div className={styles.specials}>
                    {specialsItems.map((item, index) => (
                        <Link to={item.link} key={index}>
                            <ul className={styles.specialList}>
                                <li className={styles.specialListItem}>{item.title}</li>
                                <li className={styles.specialListItem}>{item.icon}</li>
                                <li className={styles.specialListItem}>{item.description}</li>
                            </ul>
                        </Link>
                    ))}
                </div>
            </div>
        </section>
    );
};

export default Displays;
