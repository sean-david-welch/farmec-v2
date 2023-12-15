import { Component } from 'solid-js';
import utils from '../../styles/Info.module.css';

const Info: Component = () => {
  return (
    <div class={utils.info}>
      <div class={utils.infoSection}>
        <h1 class={utils.subHeading}>Business Information:</h1>
        <div class={utils.infoItem}>
          Opening Hours:
          <br />
          <span class={utils.infoItemText}>Monday - Friday: 9am - 5:30pm</span>
        </div>
        <div class={utils.infoItem}>
          Telephone:
          <br />
          <span class={utils.infoItemText}>
            <a href="tel:01 825 9289">01 825 9289</a>
          </span>
        </div>
        <div class={utils.infoItem}>
          International:
          <br />
          <span class={utils.infoItemText}>
            <a href="tel:+353 1 825 9289">+353 1 825 9289</a>
          </span>
        </div>
        <div class={utils.infoItem}>
          Email:
          <br />
          <span class={utils.infoItemText}>Info@farmec.ie</span>
        </div>
        <div class={utils.infoItem}>
          Address:
          <br />
          <span class={utils.infoItemText}>Clonross, Drumree, Co. Meath, A85PK30</span>
        </div>
      </div>
    </div>
  );
};

export default Info;
