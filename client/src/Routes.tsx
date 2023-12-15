import { lazy } from 'solid-js';
import { Router, Route } from '@solidjs/router';

const Home = lazy(() => import('./routes/Home'));
const About = lazy(() => import('./routes/About'));

function AppRoutes() {
  return (
    <Router>
      <Route path="/" component={Home} />
      <Route path="/about" component={About} />
    </Router>
  );
}

export default AppRoutes;
