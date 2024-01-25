import styles from '../../styles/Info.module.css';
import utils from '../../styles/Utils.module.css';

const icons = {
    faTractor: '../icons/tractor.svg',
    faToolbox: '../icons/toolbox.svg',
    faGears: '../icons/gears.svg',
    faUserPLus: '../icons/user-plus.svg',
};

const specialsItems = [
    {
        title: 'Quality Stock',
        description:
            'Farmec is committed to the importation and distribution of only quality brands of unique farm machinery. We guarantee that all our suppliers are committed to providing farmers with durable and superior stock',
        icon: icons.faTractor,
        link: '/about',
    },
    {
        title: 'Assemably',
        description:
            'Farmec have a team of qualified and experienced staff that ensure abundant care is taken during the assembly process; we make sure that a quality supply chain is maintained from manufacturer to customer',
        icon: icons.faToolbox,
        link: '/about',
    },
    {
        title: 'Spare Parts',
        description:
            'Farmec offers a diverse and complete range of spare parts for all its machinery. Quality stock control and industry expertise ensures parts finds their way to you efficiently',
        icon: icons.faGears,
        link: '/about',
    },
    {
        title: 'Customer Service',
        description:
            'Farmec is a family run company, we make sure we extend the ethos of a small community to our customers. We build established relationships with our dealers that provide them and the farmers with extensive guidance',
        icon: icons.faUserPLus,
        link: '/about',
    },
];

const SpecialsComponents: React.FC = () => {
    return (
        <div className={styles.infoSection}>
            <h1 className={utils.sectionHeading}>Farmec At A Glance:</h1>
            <p className={utils.subHeading}>This is a Quick Look at what Separates us from our Competitors</p>
            <div className={styles.specials}>
                {specialsItems.map(item => (
                    <a href={item.link}>
                        <ul className={styles.specialsList}>
                            <li className={styles.specialsListItem}>{item.title}</li>
                            <li className={styles.specialsListItem}>
                                <img src={item.icon} alt="icon" />
                            </li>
                            <li className={styles.specialsListItem}>{item.description}</li>
                        </ul>
                    </a>
                ))}
            </div>
        </div>
    );
};

export default SpecialsComponents;
