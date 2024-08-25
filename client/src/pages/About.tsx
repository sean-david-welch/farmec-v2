import utils from '../styles/Utils.module.css';

import Timelines from '../templates/Timeline';

import {useUserStore} from '../lib/store';
import {Employees} from '../templates/Employees';


import {Helmet} from 'react-helmet';
import {FC} from "react";
import TimelineForm from '../forms/TimelineForm';

const About: FC = () => {
    const {isAdmin} = useUserStore();

    return (
        <>
            <Helmet>
                <title>About Us - Farmec Ireland</title>
                <meta name="description" content="Learn more about Farmec's staff and history. Meet our team and explore our 30-year journey as a family-run business providing high-quality machinery." />

                <meta property="og:title" content="About Us - Farmec Ireland" />
                <meta property="og:description" content="Discover Farmec's staff, history, and vision for the future." />
                <meta property="og:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp" />
                <meta property="og:url" content="https://www.farmec.ie/about" />
                <meta property="og:type" content="website" />

                <meta name="twitter:card" content="summary_large_image" />
                <meta name="twitter:title" content="About Us - Farmec Ireland" />
                <meta name="twitter:description" content="Learn more about Farmec's staff and history. Meet our team and explore our 30-year journey as a family-run business providing high-quality machinery." />
                <meta name="twitter:image" content="https://www.farmec.ie/farmec_images/Suppliers/sip1250.webp" />
                <link rel="canonical" href="https://www.farmec.ie/about" />
            </Helmet>

            <section id="about">
                <h1 className={utils.sectionHeading}>Our Team</h1>
                <p className={utils.subHeading}>Meet our staff and management team</p>
                <Employees/>
                <div className={utils.aboutHeading}>
                    <h1 className={utils.sectionHeading}>Our Story</h1>
                    <p className={utils.subHeading}>
                        Farmec is a family run and owned business founded nearly 30 years ago. Farmec is committed to
                        the
                        importation and distribution of high quality machinery tailored to meet the needs of
                        agricultural
                        and amenity dealers and ultimately farmers, contractors and golf courses. Farmec commits to
                        supplying a reliable, professional and customer orientated sales, technical support and parts
                        service at value for money prices. Our vision is to be the preferred source of world class
                        machinery
                        for the Irish farmer to help them to grow their farm enterprises and continue to make a valuable
                        and
                        increased contribution to the Irish Economy.
                    </p>
                </div>
                {isAdmin && <TimelineForm/>}
                <Timelines/>
            </section>

        </>
    );
};

export default About;
