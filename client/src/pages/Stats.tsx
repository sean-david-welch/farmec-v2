import { Component, For } from 'solid-js';
import styles from '../../styles/Info.module.css';
import utils from '../../styles/Utils.module.css';

const icons = {
  faUsers: '../icons/users.svg',
  faBusinessTime: '../icons/business-time.svg',
  faHandshake: '../icons/handshake.svg',
  faWrench: '../icons/wrench.svg',
};

const statsItems = [
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

const StatsComponent: Component = () => {
  return (
    <div class={styles.infoSection}>
      <h1 class={utils.sectionHeading}>Farmec At A Glance:</h1>
      <p class={utils.subHeading}>This is a Quick Look at what Separates us from our Competitors</p>
      <div class={styles.stats}>
        <For each={statsItems}>
          {item => (
            <a href={item.link}>
              <ul class={styles.statList}>
                <li class={styles.statListItem}>{item.title}</li>
                <li class={styles.statListItem}>
                  <img src={item.icon} alt="icon" />
                </li>
                <li class={styles.statListItem}>{item.description}</li>
              </ul>
            </a>
          )}
        </For>
      </div>
    </div>
  );
};

export default StatsComponent;
