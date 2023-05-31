package lib

const (
	XApiKey      = "REPLACE_WITH_YOUR_KEY"
	XProductCode = "REPLACE_WITH_PRODUCT_CODE"
)

const GatewayHost = "https://gateway.snowpal.com/"

const (
	RouteAttributesGetDisplayableAttributes       = "app/resource/attributes"
	RouteAttributesUpdateKeyDisplayAttributes     = "keys/%s/attributes"
	RouteAttributesUpdateProjectDisplayAttributes = "projects/%s/attributes?keyId=%s"
	RouteAttributesUpdateCardDisplayAttributes    = "cards/%s/attributes?keyId=%s&projectId=%s"
)

const (
	RouteProjectsGetProjects                             = "keys/%s/projects?filter=%s&batchIndex=%s&aclWriteOrHigher=%s"
	RouteProjectsAddProject                              = "keys/%s/projects"
	RouteProjectsGetProjectsLinkedToCard                 = "cards/:id/linked-to/projects?keyId=%s"
	RouteProjectsAddProjectBasedOnTemplate               = "keys/%s/projects/by-template?templateId=%s&excludeCards=%s&excludeTasks=%s"
	RouteProjectsLinkProjectToKey                        = "keys/%s/projects/%s/link"
	RouteProjectsUnlinkProjectFromKey                    = "keys/%s/projects/%s/unlink"
	RouteProjectsGetProjectsAvailableToBeLinkedToThisKey = "keys/%s/projects/available-to-link"
	RouteProjectsGetProject                              = "projects/%s?keyId=%s"
	RouteProjectsUpdateProject                           = "projects/%s?keyId=%s"
	RouteProjectsAddProjectTypeToProject                 = "projects/%s/project-types/%s?keyId=%s"
	RouteProjectsDeleteProjectTypeFromProject            = "projects/%s/project-types?keyId=%s"
	RouteProjectsAddScaleToProject                       = "projects/%s/scales/%s?keyId=%s"
	RouteProjectsDeleteScaleFromProject                  = "projects/%s/scales?keyId=%s"
	RouteProjectsUpdateProjectScaleValue                 = "projects/%s/scale-value?keyId=%s"
	RouteProjectsUpdateProjectDescription                = "projects/%s/description?keyId=%s"
	RouteProjectsArchiveProject                          = "projects/%s/archive?keyId=%s"
	RouteProjectsUnarchiveProject                        = "projects/%s/unarchive?keyId=%s"
	RouteProjectsGetArchivedProjects                     = "projects/archived?keyId=%s&batchIndex=%s"
	RouteProjectsBulkArchiveProjects                     = "projects/archive?keyId=%s"
	RouteProjectsAllowArchivalOfProject                  = "projects/%s/allow-archival?keyId=%s"
	RouteProjectsCopyProject                             = "projects/%s/copy?keyId=%s&allTasks=%s&cardIds=%s&allCards=%s&allChecklists=%s&targetKeyId=%s"
	RouteProjectsMoveProject                             = "projects/%s/move?keyId=%s&targetKeyId=%s"
)

const (
	RouteProjectsGetProjectAttachments   = "projects/%s/attachments?keyId=%s"
	RouteProjectsAddProjectAttachment    = "projects/%s/attachments?keyId=%s"
	RouteProjectsRenameProjectAttachment = "project-attachments/%s/rename?keyId=%s&projectId=%s"
	RouteProjectsDeleteProjectAttachment = "project-attachments/%s?keyId=%s&projectId=%s"
)

const (
	RouteProjectsGetLinkedCards              = "charts/keys/%s/projects/%s/linked-resources"
	RouteProjectsGetScaleValuesForScale      = "charts/keys/%s/projects/%s/scales/%s/grades"
	RouteProjectsGetTaskStatusForProject     = "charts/keys/%s/projects/%s/task-status"
	RouteProjectsGetProjectGradesForStudents = "projects/%s/students/all/grades?keyId=%s"
)

const (
	RouteProjectsGetProjectChecklists         = "projects/%s/checklists?keyId=%s"
	RouteProjectsAddProjectChecklist          = "projects/%s/checklists?keyId=%s"
	RouteProjectsReorderProjectChecklists     = "projects/%s/checklists/reorder?keyId=%s"
	RouteProjectsDeleteProjectChecklist       = "projects/%s/checklists/%s?keyId=%s"
	RouteProjectsRenameProjectChecklist       = "projects/%s/checklists/%s?keyId=%s"
	RouteProjectsAddProjectChecklistItem      = "projects/%s/checklists/%s/checklist-items?keyId=%s"
	RouteProjectsUpdateProjectChecklistItem   = "projects/%s/checklists/%s/checklist-items/%s?keyId=%s"
	RouteProjectsDeleteProjectChecklistItem   = "projects/%s/checklists/%s/checklist-items/%s?keyId=%s"
	RouteProjectsReorderProjectChecklistItems = "projects/%s/checklists/%s/checklist-items/reorder?keyId=%s"
)

const (
	RouteProjectsGetProjectComments   = "projects/%s/comments?keyId=%s"
	RouteProjectsAddProjectComment    = "projects/%s/comments?keyId=%s"
	RouteProjectsUpdateProjectComment = "project-comments/%s?keyId=%s&projectId=%s"
	RouteProjectsDeleteProjectComment = "project-comments/%s?keyId=%s&projectId=%s"
)

const (
	RouteProjectsGetProjectNotes   = "projects/%s/notes?keyId=%s"
	RouteProjectsAddProjectNote    = "projects/%s/notes?keyId=%s"
	RouteProjectsUpdateProjectNote = "project-notes/%s?keyId=%s&projectId=%s"
	RouteProjectsDeleteProjectNote = "project-notes/%s?keyId=%s&projectId=%s"
)

const (
	RouteProjectsGetProjectTasks     = "projects/%s/tasks?keyId=%s"
	RouteProjectsAddProjectTask      = "projects/%s/tasks?keyId=%s"
	RouteProjectsUpdateProjectTask   = "project-tasks/%s?keyId=%s&projectId=%s"
	RouteProjectsDeleteProjectTask   = "project-tasks/%s?keyId=%s&projectId=%s"
	RouteProjectsAssignProjectTask   = "project-tasks/%s/assign?keyId=%s&projectId=%s"
	RouteProjectsUnassignProjectTask = "project-tasks/%s/unassign?keyId=%s&projectId=%s"
	RouteProjectsReorderProjectTasks = "projects/%s/tasks/reorder?keyId=%s"
)

const (
	RouteCardsGetCards                                 = "projects/%s/cards?batchIndex=%s&keyId=%s"
	RouteCardsAddCard                                  = "projects/%s/cards?keyId=%s"
	RouteCardsAddCardBasedOnTemplate                   = "projects/%s/cards/by-template?keyId=%s&templateId=%s&excludeTasks=%s"
	RouteCardsLinkCardToProject                        = "projects/%s/cards/%s/link?keyId=%s"
	RouteCardsUnlinkCardFromProject                    = "projects/%s/cards/%s/unlink?keyId=%s"
	RouteCardsGetCard                                  = "cards/%s?keyId=%s&projectId=%s"
	RouteCardsUpdateCard                               = "cards/%s?keyId=%s&projectId=%s"
	RouteCardsUpdateCardCompletionStatus               = "cards/%s/by-completion-status?keyId=%s&projectId=%s"
	RouteCardsAddCardTypeToCard                        = "cards/%s/card-types/%s?keyId=%s&projectId=%s"
	RouteCardsDeleteCardTypeFromCard                   = "cards/%s/card-types?keyId=%s&projectId=%s"
	RouteCardsAddScaleToCard                           = "cards/%s/scales/%s?keyId=%s&projectId=%s"
	RouteCardsDeleteScaleFromCard                      = "cards/%s/scales?keyId=%s&projectId=%s"
	RouteCardsUpdateCardScaleValue                     = "cards/%s/scale-value?keyId=%s&projectId=%s"
	RouteCardsArchiveCard                              = "cards/%s/archive?keyId=%s&project=%s"
	RouteCardsGetArchivedCards                         = "cards/archived?batchIndex=%s&keyId=%s&projectId=%s"
	RouteCardsGetCardsAvailableToBeLinkedToThisProject = "projects/%s/cards/available-to-link?keyId=%s"
	RouteCardsUnarchiveCard                            = "cards/%s/unarchive?keyId=%s&projectId=%s"
	RouteCardsBulkArchiveCards                         = "cards/archive?keyId=%s&projectId=%s"
	RouteCardsUpdateCardDescription                    = "cards/%s/description?keyId=%s&projectId=%s"
	RouteCardsAllowArchivalOfCard                      = "cards/%s/allow-archival?keyId=%s&projectId=%s"
	RouteCardsCopyCard                                 = "cards/%s/copy?keyId=%s&projectId=%s&allTasks=%s&allChecklists=%s&targetKeyId=%s&targetProjectId=%s"
	RouteCardsMoveCard                                 = "cards/%s/move?keyId=%s&projectId=%s&targetKeyId=%s&targetProjectId=%s"
)

const (
	RouteCardsGetCardAttachments   = "cards/%s/attachments?keyId=%s&projectId=%s"
	RouteCardsAddCardAttachment    = "cards/%s/attachments?keyId=%s&projectId=%s"
	RouteCardsRenameCardAttachment = "card-attachments/%s/rename?keyId=%s&projectId=%s&cardId=%s"
	RouteCardsDeleteCardAttachment = "card-attachments/%s?keyId=%s&projectId=%s&cardId=%s"
)

const (
	RouteCardsGetCardTasksForCharts = "charts/cards/%s/tasks?keyId=%s&projectId=%s"
)

const (
	RouteCardsGetCardChecklists         = "cards/%s/checklists?keyId=%s&projectId=%s"
	RouteCardsAddCardChecklist          = "cards/%s/checklists?keyId=%s&projectId=%s"
	RouteCardsReorderCardChecklists     = "cards/%s/checklists/reorder?keyId=%s&projectId=%s"
	RouteCardsDeleteCardChecklist       = "cards/%s/checklists/%s?keyId=%s&projectId=%s"
	RouteCardsRenameCardChecklist       = "cards/%s/checklists/%s?keyId=%s&projectId=%s"
	RouteCardsAddCardChecklistItem      = "cards/%s/checklists/%s/checklist-items?keyId=%s&projectId=%s"
	RouteCardsUpdateCardChecklistItem   = "cards/%s/checklists/%s/checklist-items/%s?keyId=%s&projectId=%s"
	RouteCardsDeleteCardChecklistItem   = "cards/%s/checklists/%s/checklist-items/%s?keyId=%s&projectId=%s"
	RouteCardsReorderCardChecklistItems = "cards/%s/checklists/%s/checklist-items/reorder?keyId=%s&projectId=%s"
)

const (
	RouteCardsGetCardComments   = "cards/%s/comments?keyId=%s&projectId=%s"
	RouteCardsAddCardComment    = "cards/%s/comments?keyId=%s&projectId=%s"
	RouteCardsUpdateCardComment = "card-comments/%s?keyId=%s&projectId=%s&cardId=%s"
	RouteCardsDeleteCardComment = "card-comments/%s?keyId=%s&projectId=%s&cardId=%s"
)

const (
	RouteCardsGetCardNotes   = "cards/%s/notes?keyId=%s&projectId=%s"
	RouteCardsAddCardNote    = "cards/%s/notes?keyId=%s&projectId=%s"
	RouteCardsUpdateCardNote = "card-notes/%s?keyId=%s&projectId=%s&cardId=%s"
	RouteCardsDeleteCardNote = "card-notes/%s?keyId=%s&projectId=%s&cardId=%s"
)

const (
	RouteCardsGetCardTasks     = "cards/%s/tasks?keyId=%s&projectId=%s"
	RouteCardsAddCardTask      = "cards/%s/tasks?keyId=%s&projectId=%s"
	RouteCardsUpdateCardTask   = "card-tasks/%s?keyId=%s&projectId=%s&cardId=%s"
	RouteCardsDeleteCardTask   = "card-tasks/%s?keyId=%s&projectId=%s&cardId=%s"
	RouteCardsAssignCardTask   = "card-tasks/%s/assign?keyId=%s&projectId=%s&cardId=%s"
	RouteCardsUnassignCardTask = "card-tasks/%s/unassign?keyId=%s&projectId=%s&cardId=%s"
	RouteCardsReorderCardTasks = "cards/%s/tasks/reorder?keyId=%s&projectId=%s"
)

const (
	RouteProjectTypesGetProjectTypes             = "project-types?includeCounts=%s"
	RouteProjectTypesAddProjectType              = "project-types"
	RouteProjectTypesUpdateProjectType           = "project-types/%s"
	RouteProjectTypesDeleteProjectType           = "project-types/%s"
	RouteProjectTypesGetProjectsUsingProjectType = "project-types/%s/projects"
)

const (
	RouteCollaborationGetProjectCollaborators                    = "projects/%s/acl?keyId=%s"
	RouteCollaborationUpdateProjectAcl                           = "projects/%s/users/%s/acl?keyId=%s"
	RouteCollaborationUnshareProjectFromCollaborator             = "projects/%s/users/%s/unshare?keyId=%s"
	RouteCollaborationShareProjectWithCollaborator               = "projects/%s/users/%s/share?keyId=%s"
	RouteCollaborationShareProjectWithCollaboratorAlongWithCards = "projects/%s/users/%s/share/with/cards?keyId=%s"
	RouteCollaborationGetUsersThisProjectCanBeSharedWith         = "search/projects/%s/shareable/users?keyId=%s&token=%s"
	RouteCollaborationBulkShareProjectsWithCollaborators         = "projects/users/%s/share?keyId=%s"
	RouteCollaborationLeaveProject                               = "projects/%s/leave?keyId=%s"
)

const (
	RouteCommentsGetRecentComments = "comments"
)

const (
	RouteConversationsGetUnreadConversationsCount         = "conversations/unread-status"
	RouteConversationsGetUserConversations                = "conversations"
	RouteConversationsAddPrivateOrGroupConversation       = "conversations"
	RouteConversationsGetConversationForGivenUsernames    = "conversations/by-usernames?userNames=%s"
	RouteConversationsSendMessageToAnExistingConversation = "conversations/%s/messages"
	RouteConversationsGetConversation                     = "conversations/%s"
	RouteConversationsDeleteConversation                  = "conversations/%s"
	RouteConversationsLeaveConversation                   = "conversations/%s/leave"
	RouteConversationsArchiveConversation                 = "conversations/%s/archive"
)

const (
	RouteDashboardGetDashboardDetails                 = "dashboard/combined-responses"
	RouteDashboardGetRecentlyModifiedProjectsAndCards = "dashboard/recently-modified"
	RouteDashboardGetUnreadCount                      = "dashboard/unread-count"
	RouteDashboardGetRecentlyModifiedKeys             = "dashboard/recently-modified/keys"
	RouteDashboardGetCardsAndTasksDueShortly          = "dashboard/due-shortly/cards-and-tasks"
	RouteDashboardGetProjectsDueShortly               = "dashboard/due-shortly/projects"
	RouteDashboardGetUnreadNotifications              = "dashboard/notifications/unread-status"
	RouteDashboardGetUnreadConversations              = "dashboard/conversations/unread-status"
)

const (
	RouteDashboardGetUserKeysProjectsAndCards           = "charts/dashboard/keys-projects-cards"
	RouteDashboardGetSystemKeysProjectsAndCards         = "charts/dashboard/system-keys"
	RouteDashboardGetFilteredUserKeysProjectsAndCards   = "charts/dashboard/keys/filters"
	RouteDashboardGetFilteredSystemKeysProjectsAndCards = "charts/dashboard/system-keys/filters"
	RouteDashboardGetProjectsBasedOnProjectTypes        = "charts/dashboard/project-types"
	RouteDashboardGetCardsBasedOnCardTypes              = "charts/dashboard/card-types"
	RouteDashboardGetProjectsAndCardsBasedOnScales      = "charts/dashboard/scales"
	RouteDashboardGetTasksByStatus                      = "charts/dashboard/task-status"
)

const (
	RouteFavoritesGetFavorites         = "favorites"
	RouteFavoritesAddKeyAsFavorite     = "favorites/keys/%s"
	RouteFavoritesAddProjectAsFavorite = "favorites/projects/%s?keyId=%s"
	RouteFavoritesAddCardAsFavorite    = "favorites/cards/%s?keyId=%s&projectId=%s"
	RouteFavoritesDeleteFavorite       = "favorites/%s"
)

const (
	RouteFollowersAddUserToFollowUsList = "app/users/follow-us"
	RouteFollowersGetFollowers          = "app/followers"
)

const (
	RouteKeysGetKeys                = "keys?batchIndex=%s"
	RouteKeysAddKey                 = "keys"
	RouteKeysAddKeyBasedOnTemplate  = "keys/by-template?templateId=%s&excludeProjects=%s&excludeCards=%s&excludeTasks=%s"
	RouteKeysGetKey                 = "keys/%s"
	RouteKeysUpdateKey              = "keys/%s"
	RouteKeysGetArchivedKeys        = "keys/archived"
	RouteKeysGetKeysLinkedToCard    = "cards/%s/linked-to/keys?keyId=%s"
	RouteKeysGetKeysLinkedToProject = "projects/%s/linked-to/keys?keyId=%s"
	RouteKeysGetKeysFilteredByType  = "keys/filtered/by-type?keyType=%s"
	RouteKeysBulkArchiveKeys        = "keys/archive?keyIds=%s"
	RouteKeysArchiveKey             = "keys/%s/archive"
	RouteKeysUnarchiveKey           = "keys/%s/unarchive"
	RouteKeysUpdateKeyDescription   = "keys/%s/description"
)

const (
	RouteKeysGetProjectsAndCardsAssociatedWithKey           = "charts/keys/%s/projects-cards"
	RouteKeysGetFilteredUserKeysProjectsAndCardsForGivenKey = "charts/keys/%s/filters"
	RouteKeysGetProjectTypesAndProjectsBasedOnThemInKey     = "charts/keys/%s/project-types"
	RouteKeysGetCardsBasedOnCardTypesInKey                  = "charts/keys/%s/card-types"
	RouteKeysGetScalesAlongWithProjectsAndCardsBasedOnThem  = "charts/keys/%s/scales"
	RouteKeysGetLinkedResources                             = "charts/keys/%s/linked-resources"
	RouteKeysGetProjectScaleValues                          = "charts/keys/%s/scales/%s/scale-values"
	RouteKeysGetTaskStatus                                  = "charts/keys/%s/task-status"
)

const (
	RouteKeysGetKeyChecklists         = "keys/%s/checklists"
	RouteKeysAddKeyChecklist          = "keys/%s/checklists"
	RouteKeysReorderKeyChecklists     = "keys/%s/checklists/reorder"
	RouteKeysDeleteKeyChecklist       = "keys/%s/checklists/%s"
	RouteKeysRenameKeyChecklist       = "keys/%s/checklists/%s"
	RouteKeysAddKeyChecklistItem      = "keys/%s/checklists/%s/checklist-items"
	RouteKeysUpdateKeyChecklistItem   = "keys/%s/checklists/%s/checklist-items/%s"
	RouteKeysDeleteKeyChecklistItem   = "keys/%s/checklists/%s/checklist-items/%s"
	RouteKeysReorderKeyChecklistItems = "keys/%s/checklists/%s/checklist-items/reorder"
)

const (
	RouteKeysGetKeyNotes   = "keys/%s/notes"
	RouteKeysAddKeyNote    = "keys/%s/notes"
	RouteKeysUpdateKeyNote = "key-notes/%s?keyId=%s"
	RouteKeysDeleteKeyNote = "key-notes/%s?keyId=%s"
)

const (
	RouteKeysGetKeyTasks     = "keys/%s/tasks"
	RouteKeysAddKeyTask      = "keys/%s/tasks"
	RouteKeysUpdateKeyTask   = "key-tasks/%s?keyId=%s"
	RouteKeysDeleteKeyTask   = "key-tasks/%s?keyId=%s"
	RouteKeysReorderKeyTasks = "keys/%s/tasks/reorder"
)

const (
	RouteNotificationsGetNotifications              = "notifications"
	RouteNotificationsGetUnreadNotifications        = "notifications/unread"
	RouteNotificationsGetUnreadNotificationCount    = "notifications/unread/count"
	RouteNotificationsMarkNotificationAsRead        = "notifications/%s/read"
	RouteNotificationsMarkNotificationsAsReadInBulk = "notifications/read"
)

const (
	RouteCardTypesGetCardTypes          = "card-types?includeCounts=%s"
	RouteCardTypesAddCardType           = "card-types"
	RouteCardTypesUpdateCardType        = "card-types/%s"
	RouteCardTypesDeleteCardType        = "card-types/%s"
	RouteCardTypesGetCardsUsingCardType = "card-types/%s/cards"
)

const (
	RouteProfileGetUsersProfile              = "profiles"
	RouteProfileUpdateUsersProfile           = "profiles"
	RouteProfileUpdateUsername               = "profiles/username/%s"
	RouteProfileBlockUserFromSendingMessages = "users/%s/block"
	RouteProfileUnblockUser                  = "users/%s/unblock"
)

const (
	RouteProjectKeysAddACard               = "projects/%s/cards?keyId=%s&projectListId=%s"
	RouteProjectKeysAddCardBasedOnTemplate = "projects/%s/cards/by-template?keyId=%s&projectListId=%s&templateId=%s&excludeTasks=%s"
	RouteProjectKeysLinkCardToProject      = "projects/%s/cards/%s/link?keyId=%s&projectListId=%s"
	RouteProjectKeysReorderCards           = "projects/%s/cards/reorder?keyId=%s"
	RouteProjectKeysCopyCard               = "cards/%s/copy?keyId=%s&projectId=%s&targetKeyId=%s&targetProjectId=%s&targetProjectListId=%s"
	RouteProjectKeysMoveCard               = "cards/%s/move?keyId=%s&projectId=%s&targetKeyId=%s&targetProjectId=%s&targetProjectListId=%s"
	RouteProjectKeysCopyProject            = "projects/%s/cards/copy?keyId=%s&targetKeyId=%s&allCards=%s&allTasks=%s"
	RouteProjectKeysAssignCard             = "cards/%s/assign?keyId=%s&projectId=%s"
	RouteProjectKeysUnassignCard           = "cards/%s/unassign?keyId=%s&projectId=%s"
)

const (
	RouteProjectKeysAddProjectList             = "projects/%s/project-lists?keyId=%s"
	RouteProjectKeysGetProjectLists            = "projects/%s/project-lists?keyId=%s"
	RouteProjectKeysCopyCardsInProjectList     = "project-lists/%s/cards/copy-all?keyId=%s&projectId=%s&targetKeyId=%s&targetProjectId=%s&targetProjectListId=%s&allCards=%s&allTasks=%s"
	RouteProjectKeysBulkCopyCardsInProjectList = "project-lists/%s/cards/copy?keyId=%s&projectId=%s&targetKeyId=%s&targetProjectId=%s&targetProjectListId=%s&allTasks=%s&cardIds=%s"
	RouteProjectKeysMoveCardsInProjectList     = "project-lists/%s/cards/move-all?keyId=%s&projectId=%s&targetKeyId=%s&targetProjectId=%s&targetProjectListId=%s&allCards=%s"
	RouteProjectKeysBulkMoveCardsInProjectList = "project-lists/%s/cards/move?keyId=%s&projectId=%s&targetKeyId=%s&targetProjectId=%s&targetProjectListId=%s&cardIds=%s"
	RouteProjectKeysMoveProjectList            = "projects/%s/project-lists/%s/move?keyId=%s&targetKeyId=%s&targetProjectId=%s&targetPosition=%s"
	RouteProjectKeysGetProjectList             = "projects/%s/project-lists/%s?keyId=%s"
	RouteProjectKeysRenameProjectList          = "projects/%s/project-lists/%s?keyId=%s"
	RouteProjectKeysReorderProjectList         = "projects/%s/project-lists/reorder?keyId=%s"
	RouteProjectKeysArchiveProjectList         = "projects/%s/project-lists/%s/archive?keyId=%s"
)

const (
	RouteRegistrationRegisterNewUserByEmail = "app/users/sign-up"
	RouteRegistrationSignInByEmail          = "app/users/sign-in"
	RouteRegistrationResetPassword          = "app/users/reset-password"
	RouteRegistrationActivateUser           = "app/user-verified/%s"
)

// TODO(Anish,3,03/23/23): As the endpoint is same for key card & project card. We need to create these constants to
// support both with the same endpoint with different query string params.
const (
	RouteRelationsGetRelationsForKeyMatchingSearchToken     = "search/relations?token=%s&currentKeyId=%s"
	RouteRelationsGetRelationsForProjectMatchingSearchToken = "search/relations?token=%s&currentProjectId=%s"
	RouteRelationsGetRelationsForCardMatchingSearchToken    = "search/relations?token=%s&currentCardId=%s&keyId=%s&projectId=%s"

	RouteRelationsRelateCardToKey         = "keys/%s/cards/%s/relate?targetKeyId=%s&targetProjectId=%s"
	RouteRelationsUnrelateCardFromKey     = "keys/%s/cards/%s/unrelate?targetKeyId=%s&targetProjectId=%s"
	RouteRelationsRelateCardToProject     = "projects/%s/cards/%s/relate?targetKeyId=%s&targetProjectId=%s"
	RouteRelationsUnrelateCardFromProject = "projects/%s/cards/%s/unrelate?targetKeyId=%s&targetProjectId=%s"
	RouteRelationsRelateCardToCard        = "cards/%s/cards/%s/relate?sourceKeyId=%s&sourceProjectId=%s&targetKeyId=%s&targetProjectId=%s"
	RouteRelationsUnrelateCardFromCard    = "cards/%s/cards/%s/unrelate?sourceKeyId=%s&sourceProjectId=%s&targetKeyId=%s&targetProjectId=%s"
)

const (
	RouteRelationsGetRelationsForKey         = "keys/%s/relations"
	RouteRelationsGetRelationsForProject     = "projects/%s/relations?keyId=%s"
	RouteRelationsGetRelationsForCard        = "cards/%s/relations?keyId=%s&projectId=%s"
	RouteRelationsRelateKeyToKey             = "keys/%s/keys/%s/relate"
	RouteRelationsUnrelateKeyFromKey         = "keys/%s/keys/%s/unrelate"
	RouteRelationsRelateProjectToKey         = "keys/%s/projects/%s/relate"
	RouteRelationsUnrelateProjectFromKey     = "keys/%s/projects/%s/unrelate"
	RouteRelationsRelateProjectToProject     = "projects/%s/projects/%s/relate"
	RouteRelationsUnrelateProjectFromProject = "projects/%s/projects/%s/unrelate"
)

const (
	RouteScalesGetScales             = "scales?includeCounts=%s&excludeEmpty=%s"
	RouteScalesAddScale              = "scales"
	RouteScalesGetScale              = "scales/%s"
	RouteScalesUpdateScale           = "scales/%s"
	RouteScalesDeleteScale           = "scales/%s"
	RouteScalesGetProjectsUsingScale = "scales/%s/projects"
	RouteScalesGetCardsUsingScale    = "scales/%s/cards"
)

const (
	RouteSchedulerGetEventsInGivenWindow = "scheduler/all-events?startDate=%s&endDate=%s"
	RouteSchedulerGetEventsForGivenDay   = "scheduler/all-events/by-start-date?startDate=%s"
	RouteSchedulerGetStandaloneEvents    = "scheduler/standalone-events"
	RouteSchedulerAddStandaloneEvent     = "scheduler/standalone-events"
	RouteSchedulerUpdateStandaloneEvent  = "scheduler/standalone-events/%s"
	RouteSchedulerDeleteStandaloneEvent  = "scheduler/standalone-events/%s"
)

const (
	RouteSearchSearchKeyProjectOrCardByToken = "search?token=%s"
	RouteSearchSearchUserByToken             = "search/users?token=%s"
)

const (
	RouteTeacherKeysGetAttachmentSubmissionsAsStudent = "classroom-cards/%s/submissions/attachments/as-student?keyId=%s&projectId=%s"
	RouteTeacherKeysGetCommentSubmissionsAsStudent    = "classroom-cards/%s/submissions/comments/as-student?keyId=%s&projectId=%s"
	RouteTeacherKeysGetStudentsInAProject             = "classroom-projects/%s/students?keyId=%s"
)

const (
	RouteTeacherKeysGetStudentAttachmentSubmissionsAsTeacher     = "classroom-cards/%s/submissions/attachments/as-teacher?studentId=%s&keyId=%s&projectId=%s"
	RouteTeacherKeysGetStudentCommentSubmissionsAsTeacher        = "classroom-cards/%s/submissions/comments/as-teacher?studentId=%s&keyId=%s&projectId=%s"
	RouteTeacherKeysAddAttachmentToTeacherCardAsTeacher          = "classroom-cards/%s/attachments/as-teacher?keyId=%s&projectId=%s"
	RouteTeacherKeysAddCommentToTeacherCardAsTeacher             = "classroom-cards/%s/comments/as-teacher?keyId=%s&projectId=%s"
	RouteTeacherKeysGetProjectAndCardsGradesForAStudentAsTeacher = "classroom-projects/%s/student-grades/as-teacher?studentUserId=%s&keyId=%s"
	RouteTeacherKeysPublishStudentGradesForAProject              = "classroom-projects/%s/student-grades/publish?keyId=%s"
	RouteTeacherKeysBulkPublishCardGradesForAStudent             = "classroom-cards/students/%s/grades/publish?keyId=%s&projectId=%s"
	RouteTeacherKeysBulkPublishCardGradesForStudents             = "classroom-cards/%s/students/grades/publish?keyId=%s&projectId=%s"
	RouteTeacherKeysGetProjectGradesForStudents                  = "classroom-projects/%s/students/grades?keyId=%s"
	RouteTeacherKeysGetCardGradesForStudents                     = "classroom-cards/%s/students/grades?keyId=%s&projectId=%s"
	RouteTeacherKeysAssignGradeToStudent                         = "classroom-projects/%s/student/grade?studentUserId=%s&keyId=%s"
	RouteTeacherKeysAssignCardGradeForAStudentAsTeacher          = "classroom-cards/%s/student/grade?studentUserId=%s&keyId=%s&projectId=%s"
	RouteTeacherKeysGetStudentProfile                            = "classroom/students/%s/profile?keyId=%s&projectId=%s"
)

const (
	RouteTemplatesGetKeyTemplates     = "templates/keys"
	RouteTemplatesGetProjectTemplates = "templates/projects"
	RouteTemplatesGetCardTemplates    = "templates/cards"
)

const (
	RouteUserGetUsers              = "users"
	RouteUserGetUserByUuid         = "users/uuid/%s"
	RouteUserGetUserByEmail        = "users/email/%s"
	RouteUserDeactivateUserAccount = "users/deactivate-account"
)

const (
	RouteVersionGetLatestVersion = "app/latest-version"
	RouteVersionGetAppStatus     = "app/status"
)
