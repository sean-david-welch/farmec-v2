import { Component } from 'solid-js';
import utils from '../../styles/Utils.module.css';

const Heading: Component = () => {
  return (
    <div class={utils.typewriter}>
      <h1>Importers & Distributors of Quality Agricultural Machinery</h1>

      <button class={utils.btn}>
        <a href="#Info">
          Find Out More: <img src="/icons/chevron-down.svg" alt="down" />
        </a>
      </button>
    </div>
  );
};

export default Heading;
