// Package scope contains scope definitions for every Mastodon API route.
package scope

/******************************************
 * OAuth Scopes
 * Defining what you have permission to do with the API
 * https://docs.joinmastodon.org/api/oauth-scopes/
 ******************************************/

// Public is not defined by the Mastodon API, but is used by this library
// to indicate a public route that does not require an App token.
const Public = ""

// Private is not defined by the Mastodon API, but is used by this library
// to indicate a private route that requires an App token, but does not
// require any specific scope.
const Private = "*"

// Read grants read access to every part of the account, and is the parent of every read: scope.
// https://docs.joinmastodon.org/api/oauth-scopes/#read
const Read = "read"

// ReadAccounts grants read access to profiles, preferences, endorsements, and featured tags.
const ReadAccounts = "read:accounts"

// ReadBlocks grants read access to blocked accounts and blocked domains.
const ReadBlocks = "read:blocks"

// ReadBookmarks grants read access to bookmarked statuses.
const ReadBookmarks = "read:bookmarks"

// ReadFavourites grants read access to favourited statuses.
const ReadFavourites = "read:favourites"

// ReadFilters grants read access to keyword and status filters.
const ReadFilters = "read:filters"

// ReadFollows grants read access to follows, follow requests, and followed tags.
const ReadFollows = "read:follows"

// ReadLists grants read access to lists and their members.
const ReadLists = "read:lists"

// ReadMutes grants read access to muted accounts.
const ReadMutes = "read:mutes"

// ReadNotifications grants read access to notifications.
const ReadNotifications = "read:notifications"

// ReadSearch grants access to search.
const ReadSearch = "read:search"

// ReadStatuses grants read access to statuses, timelines, conversations, polls, and markers.
const ReadStatuses = "read:statuses"

// Write grants write access to every part of the account, and is the parent of every write: scope.
// https://docs.joinmastodon.org/api/oauth-scopes/#write
const Write = "write"

// WriteAccounts grants write access to the profile, including avatar and header images.
const WriteAccounts = "write:accounts"

// WriteBlocks grants the ability to block and unblock accounts and domains.
const WriteBlocks = "write:blocks"

// WriteBookmarks grants the ability to bookmark and un-bookmark statuses.
const WriteBookmarks = "write:bookmarks"

// WriteConversations grants the ability to delete conversations and mark them read.
const WriteConversations = "write:conversations"

// WriteFavourites grants the ability to favourite statuses and react to announcements.
const WriteFavourites = "write:favourites"

// WriteFilters grants the ability to create, change, and delete filters.
const WriteFilters = "write:filters"

// WriteFollows grants the ability to follow, unfollow, and answer follow requests.
const WriteFollows = "write:follows"

// WriteLists grants the ability to create, change, and delete lists and their members.
const WriteLists = "write:lists"

// WriteMedia grants the ability to upload media attachments.
const WriteMedia = "write:media"

// WriteMutes grants the ability to mute and unmute accounts and conversations.
const WriteMutes = "write:mutes"

// WriteNotifications grants the ability to dismiss and clear notifications.
const WriteNotifications = "write:notifications"

// WriteReports grants the ability to report accounts and statuses.
const WriteReports = "write:reports"

// WriteStatuses grants the ability to publish, edit, and delete statuses.
const WriteStatuses = "write:statuses"

// Push grants access to the Web Push subscription endpoints.
// https://docs.joinmastodon.org/api/oauth-scopes/#push
const Push = "push"

// AdminRead grants read access to every moderation endpoint, and is the parent of every admin:read: scope.
// https://docs.joinmastodon.org/api/oauth-scopes/#admin
const AdminRead = "admin:read"

// AdminReadAccounts grants moderators read access to account records.
const AdminReadAccounts = "admin:read:accounts"

// AdminReadReports grants moderators read access to reports.
const AdminReadReports = "admin:read:reports"

// AdminReadDomainAllows grants moderators read access to allowed domains.
const AdminReadDomainAllows = "admin:read:domain_allows"

// AdminReadDomainBlocks grants moderators read access to blocked domains.
const AdminReadDomainBlocks = "admin:read:domain_blocks"

// AdminReadIPBlocks grants moderators read access to blocked IP ranges.
const AdminReadIPBlocks = "admin:read:ip_blocks"

// AdminReadEmailDomainBlocks grants moderators read access to blocked email domains.
const AdminReadEmailDomainBlocks = "admin:read:email_domain_blocks"

// AdminReadCanonicalEmailBlocks grants moderators read access to blocked canonical email addresses.
const AdminReadCanonicalEmailBlocks = "admin:read:canonical_email_blocks"

// AdminWrite grants write access to every moderation endpoint, and is the parent of every admin:write: scope.
const AdminWrite = "admin:write"

// AdminWriteAccounts grants moderators the ability to act on accounts.
const AdminWriteAccounts = "admin:write:accounts"

// AdminWriteReports grants moderators the ability to act on reports.
const AdminWriteReports = "admin:write:reports"

// AdminWriteDomainAllows grants moderators the ability to allow and un-allow domains.
const AdminWriteDomainAllows = "admin:write:domain_allows"

// AdminWriteDomainBlocks grants moderators the ability to block and unblock domains.
const AdminWriteDomainBlocks = "admin:write:domain_blocks"

// AdminWriteIPBlocks grants moderators the ability to block and unblock IP ranges.
const AdminWriteIPBlocks = "admin:write:ip_blocks"

// AdminWriteEmailDomainBlocks grants moderators the ability to block and unblock email domains.
const AdminWriteEmailDomainBlocks = "admin:write:email_domain_blocks"

// AdminWriteCanonicalEmailBlocks grants moderators the ability to block and unblock canonical email addresses.
const AdminWriteCanonicalEmailBlocks = "admin:write:canonical_email_blocks"

// Grants access to view a user's Collections
const ReadCollections = "read:collections"

// Grants access to create and manage a user's Collections
const WriteCollections = "write:collections"
