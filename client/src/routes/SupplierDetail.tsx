import utils from '../styles/Utils.module.css';
import styles from '../styles/Suppliers.module.css';

import Videos from '../pages/Videos';
import Machines from '../pages/Machines';

import { useParams } from '@solidjs/router';
import { Component, createResource } from 'solid-js';
import Supplier from '../types/supplier';
import Machine from '../types/machine';
import { Video } from '../types/video';

const SuppliersDetails: Component = () => {
    const params = useParams();

    const [supplier] = createResource<Supplier | undefined>(async () => {
        return fetch(`http://localhost:8080/api/suppliers/${params.id}`).then(response => response.json());
    });

    const [machines] = createResource<Machine[]>(async () => {
        return fetch(`http://localhost:8080/api/machines/${params.id}`).then(response => response.json());
    });

    const [videos] = createResource<Video[]>(async () => {
        return fetch(`http://localhost:8080/api/videos/${params.id}`).then(response => response.json());
    });

    return (
        <>
            {/* <Layout title={`Farmec | ${supplier.name}` || 'Farmec'}> */}
            <section id="supplierDetail">
                <div class={styles.supplierHeading}>
                    <h1 class={utils.sectionHeading}>{supplier.name}</h1>

                    {/* <SocialLinks
                            facebook={social_facebook}
                            twitter={social_twitter}
                            instagram={social_instagram}
                            linkedin={social_linkedin}
                            website={social_website}
                            youtube={social_youtube}
                        /> */}
                </div>

                {/* <div class={utils.index}>
                    <h1 class={utils.indexHeading}>Suppliers</h1>
                    {machines.map(link => (
                        <a href={`#${link.name}`}>
                            <h1 class="indexItem">{link.name}</h1>
                        </a>
                    ))}
                </div> */}

                <div class={styles.supplierDetail}>
                    <img
                        src={supplier()?.marketing_image ?? '/default.jpg'}
                        alt={'/dafault.jpg'}
                        class={styles.supplierImage}
                        width={750}
                        height={750}
                    />

                    <p class={styles.supplierDescription}>{supplier()?.description}</p>
                </div>

                {machines() && <Machines machines={machines() ?? []} />}
                {videos() && <Videos videos={videos() ?? []} />}
            </section>
            {/* </Layout> */}
        </>
    );
};

export default SuppliersDetails;
