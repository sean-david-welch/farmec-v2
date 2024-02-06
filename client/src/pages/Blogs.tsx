import utils from '../styles/Utils.module.css';
import styles from '../styles/Blogs.module.css';

import { Link } from 'react-router-dom';
import { Blog } from '../types/blogTypes';
import { useUserStore } from '../lib/store';
import { useGetResource } from '../hooks/genericHooks';
import BlogForm from '../forms/BlogForm';
import DeleteButton from '../components/DeleteButton';

const Blogs: React.FC = () => {
    const { isAdmin } = useUserStore();

    const blogs = useGetResource<Blog[]>('blogs');

    return (
        <section id="blog">
            <h1 className={utils.sectionHeading}>Check out our Latest Blog Posts</h1>
            <p className={utils.subHeading}> Read our latest news</p>
            {blogs.data && (
                <div className={utils.index}>
                    <h1 className={utils.indexHeading}>Suppliers</h1>
                    {blogs.data.map(link => (
                        <a key={link.title} href={`#${link.title}`}>
                            <h1 className={utils.indexItem}>{link.title}</h1>
                        </a>
                    ))}
                </div>
            )}

            {blogs.data?.map(blog => (
                <div className={styles.blogGrid} key={blog.id} id={blog.title || ''}>
                    <div className={styles.blogCard}>
                        <img src={blog.main_image || '/default.jpg'} alt={'/default.jpg'} width={400} height={400} />
                        <div className={styles.blogLink}>
                            <h1 className={utils.mainHeading}>{blog.title}</h1>
                            <p className={utils.paragraph}>{blog.subheading}</p>
                            <p className={utils.paragraph}>{blog.body}</p>
                            <button className={utils.btnForm}>
                                <Link to={`/blogs/${blog.id}`}>
                                    Read More
                                    <img src="/icons/right-bracket.svg" alt="bracket-right" />
                                </Link>
                            </button>
                        </div>
                    </div>
                    {isAdmin && blog.id && (
                        <div className={utils.optionsBtn}>
                            <BlogForm id={blog.id} blog={blog} />
                            <DeleteButton id={blog.id} resourceKey="blogs" />
                        </div>
                    )}
                </div>
            ))}
            {isAdmin && <BlogForm />}
        </section>
    );
};

export default Blogs;
