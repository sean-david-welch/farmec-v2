import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
    faTractor,
    faToolbox,
    faGears,
    faUserPlus,
    faUsers,
    faBusinessTime,
    faHandshake,
    faWrench,
} from '@fortawesome/free-solid-svg-icons';

const icons = {
    faTractor: <FontAwesomeIcon icon={faTractor} />,
    faToolbox: <FontAwesomeIcon icon={faToolbox} />,
    faGears: <FontAwesomeIcon icon={faGears} />,
    faUserPlus: <FontAwesomeIcon icon={faUserPlus} />,
    faUsers: <FontAwesomeIcon icon={faUsers} />,
    faBusinessTime: <FontAwesomeIcon icon={faBusinessTime} />,
    faHandshake: <FontAwesomeIcon icon={faHandshake} />,
    faWrench: <FontAwesomeIcon icon={faWrench} />,
};

export const specialsItems = [
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
        icon: icons.faUserPlus,
        link: '/about',
    },
];

export const statsItems = [
    {
        title: 'Large Network',
        description: '50+ Dealers Nationwide',
        icon: icons.faUsers,
        link: '/about',
    },
    {
        title: 'Experience',
        description: '25+ Years in Business',
        icon: icons.faBusinessTime,
        link: '/about',
    },
    {
        title: 'Diverse Range',
        description: '10+ Quality Suppliers',
        icon: icons.faHandshake,
        link: '/about',
    },
    {
        title: 'Committment',
        description: 'Warranty Guarentee',
        icon: icons.faWrench,
        link: '/about',
    },
];
