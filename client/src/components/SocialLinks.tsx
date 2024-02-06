import styles from '../styles/Suppliers.module.css';

import { Link } from 'react-router-dom';
import { faGlobe } from '@fortawesome/free-solid-svg-icons/faGlobe';
import { faYoutube } from '@fortawesome/free-brands-svg-icons/faYoutube';
import { faFacebook } from '@fortawesome/free-brands-svg-icons/faFacebook';
import { faTwitter } from '@fortawesome/free-brands-svg-icons/faTwitter';
import { faInstagram } from '@fortawesome/free-brands-svg-icons/faInstagram';
import { faLinkedin } from '@fortawesome/free-brands-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface SocialLinksProps {
    facebook?: string | null;
    twitter?: string | null;
    instagram?: string | null;
    linkedin?: string | null;
    website?: string | null;
    youtube?: string | null;
}

export const SocialLinks = ({
    facebook,
    twitter,
    instagram,
    linkedin,
    website,
    youtube,
}: SocialLinksProps) => {
    const facebookIcon = <FontAwesomeIcon icon={faFacebook} className={styles.icon} />;
    const twitterIcon = <FontAwesomeIcon icon={faTwitter} className={styles.icon} />;
    const instagramIcon = <FontAwesomeIcon icon={faInstagram} className={styles.icon} />;
    const linkedinIcon = <FontAwesomeIcon icon={faLinkedin} className={styles.icon} />;
    const websiteIcon = <FontAwesomeIcon icon={faGlobe} className={styles.icon} />;
    const youtubeIcon = <FontAwesomeIcon icon={faYoutube} className={styles.icon} />;

    return (
        <div className={styles.socialLinks}>
            {facebook && (
                <Link to={facebook} target="_blank" className={styles.facebookButton}>
                    {facebookIcon}
                </Link>
            )}
            {twitter && (
                <Link to={twitter} target="_blank" className={styles.twitterButton}>
                    {twitterIcon}
                </Link>
            )}
            {instagram && (
                <Link to={instagram} target="_blank" className={styles.instagramButton}>
                    {instagramIcon}
                </Link>
            )}
            {linkedin && (
                <Link to={linkedin} target="_blank" className={styles.linkedinButton}>
                    {linkedinIcon}
                </Link>
            )}
            {website && (
                <Link to={website} target="_blank" className={styles.websiteButton}>
                    {websiteIcon}
                </Link>
            )}
            {youtube && (
                <Link to={youtube} target="_blank" className={styles.youtubeButton}>
                    {youtubeIcon}
                </Link>
            )}
        </div>
    );
};
