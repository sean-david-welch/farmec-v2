import styles from '../styles/About.module.css';

import { useGetResource } from '../hooks/genericHooks';
import { Timeline as TimelineType } from '../types/aboutTypes';

export const Timeline = () => {
    const timeline = useGetResource<TimelineType[]>('timelines');

    return (
        <section id="timeline">
            <div className={styles.timeline}>
                {/* {timeline.map((event, index) => (
                    <TimelineCard
                        key={event.id}
                        event={event}
                        user={user}
                        direction={index % 2 === 0 ? 'left' : 'right'}
                    />
                ))} */}
            </div>
        </section>
    );
};
