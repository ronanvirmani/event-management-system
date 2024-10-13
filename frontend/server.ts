// server.ts
import { APP_BASE_HREF } from '@angular/common';
import { renderModule } from '@angular/platform-server';
import express from 'express';
import { fileURLToPath } from 'node:url';
import { dirname, join, resolve } from 'node:path';
import { AppServerModule } from './src/app/app.server.module'; // Correct import
import fs from 'fs'; // Import fs to read HTML files

export function app(): express.Express {
  const server = express();

  // Resolve paths
  const serverDistFolder = dirname(fileURLToPath(import.meta.url));
  const browserDistFolder = resolve(serverDistFolder, '../browser'); // Adjust if necessary
  const indexHtmlPath = join(browserDistFolder, 'index.html'); // Ensure this path is correct

  // Set the view engine to HTML
  server.set('view engine', 'html');
  server.set('views', browserDistFolder);

  // Serve static files from /browser
  server.get('*.*', express.static(browserDistFolder, {
    maxAge: '1y'
  }));

  // All regular routes use the Angular engine
  server.get('*', (req, res, next) => {
    const { originalUrl, baseUrl } = req;

    // Read the index.html file
    fs.readFile(indexHtmlPath, 'utf8', (err, document) => {
      if (err) {
        return next(err);
      }

      // Render the Angular application
      renderModule(AppServerModule, {
        url: originalUrl,
        document: document,
        extraProviders: [{ provide: APP_BASE_HREF, useValue: baseUrl }],
      })
        .then(html => res.send(html))
        .catch(error => {
          console.error('Error during server-side rendering:', error);
          next(error);
        });
    });
  });

  return server;
}

function run(): void {
  const port = process.env['PORT'] || 4000;

  // Start up the Node server
  const server = app();
  server.listen(port, () => {
    console.log(`Node Express server listening on http://localhost:${port}`);
  });
}

run();
