import styles from '../styles/About.module.css';
import utils from '../styles/Utils.module.css';

import TimelineForm from '../forms/TimelineForm';
import { Timeline } from '../types/aboutTypes';
import DeleteButton from './DeleteButton';

interface Props {
    isAdmin: boolean;
    timeline: Timeline;
}

const TimelineCard: React.FC<Props> = ({ isAdmin, timeline }) => {
    return (
        <div className={styles.timelineCard}>
            <h1 className={utils.mainHeading}>{timeline.title}</h1>
            <h1 className={utils.paragraph}>
                <img src="/icons/clock.svg" className={styles.clockIcon} />-{timeline.date}
            </h1>
            <p className={utils.paragraph}>{timeline.body}</p>
            {isAdmin && timeline.id && (
                <div className={utils.optionsBtn}>
                    <TimelineForm id={timeline.id} timeline={timeline} />
                    <DeleteButton id={timeline.id} resourceKey="timelines" />
                </div>
            )}
        </div>
    );
};

export default TimelineCard;
