'use client'; // This is crucial to mark it as a Client Component

import { usePathname } from 'next/navigation';
import Link from 'next/link';
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from '@/components/ui/breadcrumb';
import React from 'react';

// A helper function to capitalize strings
function capitalize(str: string) {
  return str.charAt(0).toUpperCase() + str.slice(1);
}

export function DynamicBreadcrumb() {
  const pathname = usePathname();
  // Split the pathname by '/' and remove any empty strings from the start
  const segments = pathname.split('/').filter(Boolean);

  // If we are on the homepage, don't show any breadcrumbs
  if (segments.length === 0) {
    return null;
  }

  return (
    <Breadcrumb>
      <BreadcrumbList>
        {/* Always add a link to the homepage */}
        <BreadcrumbItem>
          <BreadcrumbLink asChild>
            <Link href="/">Home</Link>
          </BreadcrumbLink>
        </BreadcrumbItem>

        {/* Map over the segments to create the rest of the breadcrumbs */}
        {segments.map((segment, index) => {
          const href = '/' + segments.slice(0, index + 1).join('/');
          const isLast = index === segments.length - 1;

          return (
            // Use React.Fragment to group the separator and item
            <React.Fragment key={href}>
              <BreadcrumbSeparator />
              <BreadcrumbItem>
                {isLast ? (
                  // The last item is the current page, so it's not a link
                  <BreadcrumbPage>{capitalize(segment)}</BreadcrumbPage>
                ) : (
                  // All other items are links to their respective pages
                  <BreadcrumbLink asChild>
                    <Link href={href}>{capitalize(segment)}</Link>
                  </BreadcrumbLink>
                )}
              </BreadcrumbItem>
            </React.Fragment>
          );
        })}
      </BreadcrumbList>
    </Breadcrumb>
  );
}