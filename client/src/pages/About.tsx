import utils from '../styles/Utils.module.css';

import { useUserStore } from '../lib/context';
import { Employees } from '../templates/Employees';
import { Timeline } from '../templates/Timeline';

const About: React.FC = () => {
    const { isAdmin } = useUserStore();

    return (
        <section id="about">
            <h1 className={utils.sectionHeading}>Our Team</h1>
            <p className={utils.subHeading}>Meet our staff and management team</p>
            <Employees />
            {/* {isAdmin && <AboutForm modelName="employee" />} */}
            <div className={utils.aboutHeading}>
                <h1 className={utils.sectionHeading}>Our Story</h1>
                <p className={utils.subHeading}>
                    Farmec is a family run and owned business founded nearly 30 years ago. Farmec is committed to the
                    importation and distribution of high quality machinery tailored to meet the needs of agricultural
                    and amenity dealers and ultimately farmers, contractors and golf courses. Farmec commits to
                    supplying a reliable, professional and customer orientated sales, technical support and parts
                    service at value for money prices. Our vision is to be the preferred source of world class machinery
                    for the Irish farmer to help them to grow their farm enterprises and continue to make a valuable and
                    increased contribution to the Irish Economy.
                </p>
            </div>
            <Timeline />
            {/* {isAdmin && <AboutForm modelName="timeline" />} */}
        </section>
    );
};

export default About;
