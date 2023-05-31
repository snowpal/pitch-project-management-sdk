#!/bin/sh

mkdir attributes.1
(
  # shellcheck disable=SC2164
  cd attributes.1
  touch attributes.1.1.get_displayable_attributes_of_key.go
  touch attributes.1.2.update_key_display_attributes.go
  touch attributes.1.3.update_project_display_attributes.go
  touch attributes.1.4.update_key_pod_display_attributes.go
  touch attributes.1.5.update_project_pod_display_attributes.go
)

mkdir projects.1
(
  # shellcheck disable=SC2164
  cd projects.1
  touch projects.1.1.get_projects_in_a_key.go
  touch projects.1.2.add_project.go
  touch projects.1.3.get_projects_linked_to_pod.go
  touch projects.1.4.add_project_based_on_template.go
  touch projects.1.5.link_project_to_key.go
  touch projects.1.6.unlink_project_from_key.go
  touch projects.1.7.get_all_projects_available_to_be_linked_to_this_key.go
  touch projects.1.8.get_project.go
  touch projects.1.9.update_project.go
  touch projects.1.10.add_project_type_to_project.go
  touch projects.1.11.delete_project_type_from_project.go
  touch projects.1.12.add_scale_to_project.go
  touch projects.1.13.delete_scale_from_project.go
  touch projects.1.14.add_scale_value_to_project.go
  touch projects.1.15.update_project_description.go
  touch projects.1.16.archive_project.go
  touch projects.1.17.unarchive_project.go
  touch projects.1.18.get_all_archived_projects.go
  touch projects.1.19.bulk_archive_projects.go
  touch projects.1.20.allow_archival_of_project.go
  touch projects.1.21.copy_project.go
  touch projects.1.22.move_project.go
)

mkdir projects.2.attachments
(
  # shellcheck disable=SC2164
  cd projects.2.attachments
  touch projects.2.1.get_project_attachments.go
  touch projects.2.2.add_project_attachment.go
  touch projects.2.3.rename_project_attachment.go
  touch projects.2.4.delete_project_attachment.go
)

mkdir projects.3.charts
(
  # shellcheck disable=SC2164
  cd projects.3.charts
  touch projects.3.1.get_linked_project_pods.go
  touch projects.3.2.get_scale_values_for_scale.go
  touch projects.3.3.get_task_status_for_project.go
  touch projects.3.4.get_project_grades_for_all_students.go
)

mkdir projects.4.checklists
(
  # shellcheck disable=SC2164
  cd projects.4.checklists
  touch projects.4.1.get_project_checklists.go
  touch projects.4.2.add_project_checklist.go
  touch projects.4.3.reorder_project_checklists.go
  touch projects.4.4.rename_checklist_title.go
  touch projects.4.5.delete_checklist.go
  touch projects.4.6.add_checklist_item.go
  touch projects.4.7.update_checklist_item.go
  touch projects.4.8.delete_checklist_item.go
  touch projects.4.9.reorder_checklist_items.go
)

mkdir projects.5.comments
(
  # shellcheck disable=SC2164
  cd projects.5.comments
  touch projects.5.1.get_project_comments.go
  touch projects.5.2.add_project_comment.go
  touch projects.5.3.update_project_comment.go
  touch projects.5.4.delete_project_comment.go
)

mkdir projects.6.notes
(
  # shellcheck disable=SC2164
  cd projects.6.notes
  touch projects.6.1.get_project_notes.go
  touch projects.6.2.add_project_note.go
  touch projects.6.3.update_project_note.go
  touch projects.6.4.delete_project_note.go
)

mkdir projects.7.tasks
(
  # shellcheck disable=SC2164
  cd projects.7.tasks
  touch projects.7.1.get_project_tasks.go
  touch projects.7.2.add_project_task.go
  touch projects.7.3.update_project_task.go
  touch projects.7.4.delete_project_task.go
  touch projects.7.5.assign_project_task.go
  touch projects.7.6.unassign_project_task.go
  touch projects.7.7.reorder_project_tasks.go
)

mkdir project_pods.1
(
  # shellcheck disable=SC2164
  cd project_pods.1
  touch project_pods.1.1.get_project_pods.go
  touch project_pods.1.2.add_project_pod.go
  touch project_pods.1.3.add_project_pod_based_on_template..go
  touch project_pods.1.4.link_pod_to_project.go
  touch project_pods.1.5.unlink_pod_from_project.go
  touch project_pods.1.6.get_project_pod.go
  touch project_pods.1.7.update_project_pod.go
  touch project_pods.1.8.update_project_pod_completion_status.go
  touch project_pods.1.9.add_pod_type_to_project_pod.go
  touch project_pods.1.10.remove_pod_type_from_project_pod.go
  touch project_pods.1.11.add_scale_to_project_pod.go
  touch project_pods.1.12.delete_scale_from_project_pod.go
  touch project_pods.1.13.update_project_pod_scale_value.go
  touch project_pods.1.14.archive_project_pod.go
  touch project_pods.1.15.get_archived_project_pods.go
  touch project_pods.1.16.get_pods_available_to_be_linked.go
  touch project_pods.1.17.unarchive_project_pod.go
  touch project_pods.1.18.bulk_archive_project_pods.go
  touch project_pods.1.19.update_project_pod_description.go
  touch project_pods.1.20.allow_archival_of_project_pod.go
  touch project_pods.1.21.copy_project_pod.go
  touch project_pods.1.22.move_project_pod.go
)

mkdir project_pods.2.attachments
(
  # shellcheck disable=SC2164
  cd project_pods.2.attachments
  touch project_pods.2.1.get_project_pod_attachments.go
  touch project_pods.2.2.add_project_pod_attachment.go
  touch project_pods.2.3.rename_project_pod_attachment.go
  touch project_pods.2.4.delete_project_pod_attachment.go
)

mkdir project_pods.3.charts
(
  # shellcheck disable=SC2164
  cd project_pods.3.charts
  touch project_pods.3.1.get_project_pod_tasks_for_Charts.go
  touch project_pods.3.2.get_project_pod_grades_for_all_students.go
)

mkdir project_pods.4.checklists
(
  # shellcheck disable=SC2164
  cd project_pods.4.checklists
  touch project_pods.4.1.get_project_pod_checklists.go
  touch project_pods.4.2.add_project_pod_checklist.go
  touch project_pods.4.3.reorder_project_pod_checklists.go
  touch project_pods.4.4.delete_project_pod_checklist.go
  touch project_pods.4.5.rename_project_pod_checklist.go
  touch project_pods.4.6.add_project_pod_checklist_item.go
  touch project_pods.4.7.update_project_pod_checklist_item.go
  touch project_pods.4.8.delete_project_pod_checklist_item.go
  touch project_pods.4.9.reorder_project_pod_checklist_items.go
)

mkdir project_pods.5.comments
(
  # shellcheck disable=SC2164
  cd project_pods.5.comments
  touch project_pods.5.1.get_project_pod_comments.go
  touch project_pods.5.2.add_project_pod_comment.go
  touch project_pods.5.3.update_project_pod_comment.go
  touch project_pods.5.4.delete_project_pod_comment.go
)

mkdir project_pods.6.notes
(
  # shellcheck disable=SC2164
  cd project_pods.6.notes
  touch project_pods.6.1.get_project_pod_notes.go
  touch project_pods.6.2.add_project_pod_note.go
  touch project_pods.6.3.update_project_pod_note.go
  touch project_pods.6.4.delete_project_pod_note.go
)

mkdir project_pods.7.tasks
(
  # shellcheck disable=SC2164
  cd project_pods.7.tasks
  touch project_pods.7.1.get_project_pod_tasks.go
  touch project_pods.7.2.add_project_pod_task.go
  touch project_pods.7.3.update_project_pod_task.go
  touch project_pods.7.4.delete_project_pod_task.go
  touch project_pods.7.5.assign_project_pod_task.go
  touch project_pods.7.6.unassign_project_pod_task.go
  touch project_pods.7.7.reorder_project_pod_tasks.go
)

mkdir project_types.1
(
  # shellcheck disable=SC2164
  cd project_types.1
  touch project_types.1.1.get_project_types.go
  touch project_types.1.2.add_project_type.go
  touch project_types.1.3.update_project_type.go
  touch project_types.1.4.delete_project_type.go
  touch project_types.1.5.get_projects_using_project_type.go
)

mkdir collaboration.1.projects
(
  # shellcheck disable=SC2164
  touch collaboration.1.projects.go
  touch collaboration.1.1.get_project_collaborators.go
  touch collaboration.1.2.update_collaborators_access_level.go
  touch collaboration.1.3.Unshare_project_from_collaborator.go
  touch collaboration.1.4.Share_project_with_user.go
  touch collaboration.1.5.Share_project_with_user_along_with_pods.go
  touch collaboration.1.6.get_all_users_this_project_can_be_shared_with.go
  touch collaboration.1.7.bulk_share_projects_with_collaborators.go
  touch collaboration.1.8.leave_project.go
)

mkdir collaboration.2.project_pods
(
  # shellcheck disable=SC2164
  cd collaboration.2.project_pods
  touch collaboration.2.1.get_project_pod_collaborators.go
  touch collaboration.2.2.Share_project_pod_with_user.go
  touch collaboration.2.3.Unshare_project_pod_from_user.go
  touch collaboration.2.4.bulk_share_project_pods_with_collaborators.go
  touch collaboration.2.5.get_all_users_this_project_pod_can_be_shared_with.go
  touch collaboration.2.6.update_project_pod_ACL.go
  touch collaboration.2.7.leave_project_pod.go
)

mkdir collaboration.3.key_pods
(
  # shellcheck disable=SC2164
  cd collaboration.3.key_pods
  touch collaboration.3.1.get_key_pod_collaborators.go
  touch collaboration.3.2.Share_key_pod_with_user.go
  touch collaboration.3.3.bulk_share_pods_with_collaborators.go
  touch collaboration.3.4.Unshare_key_pod_from_user.go
  touch collaboration.3.5.get_all_users_this_key_pod_can_be_shared_with.go
  touch collaboration.3.6.update_pod_acl.go
  touch collaboration.3.7.leave_pod.go
)

mkdir comments.1
(
  # shellcheck disable=SC2164
  cd comments.1
  touch comments.1.1.get_recent_comments.go
)

mkdir conversations.1
(
  # shellcheck disable=SC2164
  cd conversations.1
  touch conversations.1.1.get_unread_conversations_count.go
  touch conversations.1.2.get_user_conversations.go
  touch conversations.1.3.add_private_or_group_conversation.go
  touch conversations.1.4.get_conversation_for_given_usernames.go
  touch conversations.1.5.send_message_to_an_existing_conversation.go
  touch conversations.1.6.get_conversation.go
  touch conversations.1.7.delete_conversation.go
  touch conversations.1.8.leave_conversation.go
  touch conversations.1.9.archive_conversation.go
)

mkdir dashboard.1
(
  # shellcheck disable=SC2164
  cd dashboard.1
  touch dashboard.1.1.get_dashboard_details.go
  touch dashboard.1.2.get_recently_modified_projects_and_pods.go
  touch dashboard.1.3.get_unread_count.go
  touch dashboard.1.4.get_recently_modified_keys.go
  touch dashboard.1.5.get_pods_and_tasks_due_shortly.go
  touch dashboard.1.6.get_projects_due_shortly.go
  touch dashboard.1.7.get_unread_notifications.go
  touch dashboard.1.8.get_unread_conversations.go
)

mkdir dashboard.2.charts
(
  # shellcheck disable=SC2164
  cd dashboard.2.charts
  touch dashboard.2.1.get_user_keys.go_projects_and_pods.go
  touch dashboard.2.2.get_system_keys.go_projects_and_pods.go
  touch dashboard.2.3.get_filtered_user_keys.go_projects_and_pods..go
  touch dashboard.2.4.get_filtered_system_keys.go_projects_and_pods..go
  touch dashboard.2.5.get_projects_based_on_project_types.go
  touch dashboard.2.6.get_pods_based_on_pod_types.go
  touch dashboard.2.7.get_projects_and_pods_based_on_scales.go
  touch dashboard.2.8.get_tasks_by_status.go
)

mkdir favorites.1
(
  # shellcheck disable=SC2164
  cd favorites.1
  touch favorites.1.1.get_favorites.go
  touch favorites.1.2.add_key_as_favorite.go
  touch favorites.1.3.add_project_as_favorite.go
  touch favorites.1.4.add_pod_as_favorite.go
  touch favorites.1.5.add_project_pod_as_favorite.go
  touch favorites.1.6.delete_favorite.go
)

mkdir followers.1
(
  # shellcheck disable=SC2164
  cd followers.1
  touch followers.1.1.add_user_to_Follow_Us_list.go
  touch followers.1.2.get_followers.go
)

mkdir keys.1
(
  # shellcheck disable=SC2164
  cd keys.1
  touch keys.1.1.get_all_keys.go
  touch keys.1.2.add_key.go
  touch keys.1.3.get_key.go
  touch keys.1.4.update_key.go
  touch keys.1.5.get_archived_keys.go
  touch keys.1.6.get_keys_a_pod_is_linked_to.go
  touch keys.1.7.get_keys_a_project_is_linked_to.go
  touch keys.1.8.get_keys_filtered_by_type.go
  touch keys.1.9.bulk_archive_keys.go
  touch keys.1.10.archive_key.go
  touch keys.1.11.unarchive_key.go
  touch keys.1.12.update_key_description.go
)

mkdir keys.2.charts
(
  # shellcheck disable=SC2164
  cd keys.2.charts
  touch keys.2.1.get_projects_and_pods_associated_with_key.go
  touch keys.2.2.get_filtered_user_keys.go_projects_and_pods_for_given_key..go
  touch keys.2.3.get_project_types_and_projects_based_on_them_in_key.go
  touch keys.2.4.get_pods_based_on_pod_types_in_key.go
  touch keys.2.5.get_scales_along_with_projects_and_pods_based_on_them.go
  touch keys.2.6.get_linked_resources.go
  touch keys.2.7.get_key_pod_and_project_scale_values.go
  touch keys.2.8.get_task_status.go
)

mkdir keys.3.checklists
(
  # shellcheck disable=SC2164
  cd keys.3.checklists
  touch keys.3.1.get_key_checklists.go
  touch keys.3.2.add_key_checklist.go
  touch keys.3.3.reorder_key_checklists.go
  touch keys.3.4.rename_checklist.go
  touch keys.3.5.delete_key_checklist.go
  touch keys.3.6.add_key_checklist_item.go
  touch keys.3.7.update_key_checklist_item.go
  touch keys.3.8.delete_key_checklist_item.go
  touch keys.3.9.reorder_key_checklist_items.go
)

mkdir keys.4.notes
(
  # shellcheck disable=SC2164
  cd keys.4.Notes
  touch keys.4.1.get_key_notes.go
  touch keys.4.2.add_key_note.go
  touch keys.4.3.update_key_note.go
  touch keys.4.4.delete_key_note.go
)

mkdir keys.5.tasks
(
  # shellcheck disable=SC2164
  cd keys.5.Tasks
  touch keys.5.1.get_key_tasks.go
  touch keys.5.2.add_key_task.go
  touch keys.5.3.update_key_task.go
  touch keys.5.4.delete_key_task.go
  touch keys.5.5.reorder_key_task.go
)

mkdir key_pods.1
(
  # shellcheck disable=SC2164
  cd key_pods.1
  touch key_pods.1.1.get_key_pods.go
  touch key_pods.1.2.add_key_pod.go
  touch key_pods.1.3.get_key_pods_available_to_be_linked.go
  touch key_pods.1.4.link_key_pod_to_key.go
  touch key_pods.1.5.unlink_key_pod_from_key.go
  touch key_pods.1.6.get_key_pod.go
  touch key_pods.1.7.update_key_pod.go
  touch key_pods.1.8.update_key_pods_completion_status.go
  touch key_pods.1.9.update_key_pods_scale_value..go
  touch key_pods.1.10.add_pod_type_to_key_pod..go
  touch key_pods.1.11.remove_pod_type_from_key_pod.go
  touch key_pods.1.12.add_scale_to_pod.go
  touch key_pods.1.13.delete_scale_from_key_pod.go
  touch key_pods.1.14.archive_key_pod.go
  touch key_pods.1.15.get_archived_pods.go
  touch key_pods.1.16.unarchive_key_pod.go
  touch key_pods.1.17.bulk_archive_key_pods.go
  touch key_pods.1.18.update_key_pod_description.go
  touch key_pods.1.19.allow_archival_of_key_pod.go
  touch key_pods.1.20.copy_key_pod.go
  touch key_pods.1.21.move_key_pod.go
)

mkdir key_pods.2.Attachments
(
  # shellcheck disable=SC2164
  cd key_pods.2.Attachments
  touch key_pods.2.1.get_key_pod_attachments.go
  touch key_pods.2.2.add_attachment_to_key_pod.go
  touch key_pods.2.3.rename_key_pod_attachment.go
  touch key_pods.2.4.delete_key_pod_attachment.go
)

mkdir key_pods.3.checklists
(
  # shellcheck disable=SC2164
  touch key_pods.3.checklists
  touch key_pods.3.1.get_key_pod_checklists.go
  touch key_pods.3.2.add_key_pod_checklist.go
  touch key_pods.3.3.reorder_key_pod_checklists.go
  touch key_pods.3.4.delete_key_pod_checklist.go
  touch key_pods.3.5.rename_key_pod_checklist.go
  touch key_pods.3.6.add_key_pod_checklist_item.go
  touch key_pods.3.7.update_key_pod_checklist_item.go
  touch key_pods.3.8.delete_key_pod_checklist_item.go
  touch key_pods.3.9.reorder_key_pod_checklist_items.go
)

mkdir key_pods.4.comments
(
  # shellcheck disable=SC2164
  cd key_pods.4.comments
  touch key_pods.4.1.get_key_pod_comments.go
  touch key_pods.4.2.add_key_pod_comment.go
  touch key_pods.4.3.update_key_pod_comment.go
  touch key_pods.4.4.delete_key_pod_comment.go
)

mkdir key_pods.5.Notes
(
  # shellcheck disable=SC2164
  cd key_pods.5.Notes
  touch key_pods.5.1.get_key_pod_notes.go
  touch key_pods.5.2.add_key_pod_note.go
  touch key_pods.5.3.update_key_pod_note.go
  touch key_pods.5.4.delete_key_pod_note.go
)

mkdir key_pods.6.tasks
(
  # shellcheck disable=SC2164
  cd key_pods.6.tasks
  touch key_pods.6.1.get_key_pod_tasks.go
  touch key_pods.6.2.add_key_pod_task.go
  touch key_pods.6.3.update_key_pod_task.go
  touch key_pods.6.4.delete_key_pod_task.go
  touch key_pods.6.5.assign_key_pod_task.go
  touch key_pods.6.6.unassign_key_pod_task.go
  touch key_pods.6.7.reorder_key_pod_tasks.go
)

mkdir notifications.1
(
  # shellcheck disable=SC2164
  cd notifications.1
  touch notifications.1.1.get_all_notifications.go
  touch notifications.1.2.get_unread_notifications.go
  touch notifications.1.3.get_unread_notification_count.go
  touch notifications.1.4.mark_notification_as_read.go
  touch notifications.1.5.mark_notifications_as_read_in_bulk.go
)

mkdir pod_types.1
(
  # shellcheck disable=SC2164
  cd pod_types.1
  touch pod_types.1.1.get_pod_types.go
  touch pod_types.1.2.add_pod_type.go
  touch pod_types.1.3.update_pod_type.go
  touch pod_types.1.4.delete_pod_type.go
  touch pod_types.1.5.get_pods_using_pod_type.go
)

mkdir profile.1
(
  # shellcheck disable=SC2164
  cd profile.1
  touch profile.1.1.get_users_introduction_profile.go
  touch profile.1.2.get_users_profile.go
  touch profile.1.3.update_users_profile.go
  touch profile.1.4.update_username.go
  touch profile.1.5.projects_user_from_sending_messages.go
  touch profile.1.6.unprojects_user.go
)

mkdir project_keys.1
(
  # shellcheck disable=SC2164
  cd project_keys.1
  touch project_keys.1.1.add_a_project_pod.go
  touch project_keys.1.2.add_project_pod_based_on_template.go
  touch project_keys.1.3.link_project_pod_to_project.go
  touch project_keys.1.4.reorder_project_pods.go
  touch project_keys.1.5.copy_project_pod.go
  touch project_keys.1.6.move_project_pod.go
  touch project_keys.1.7.copy_project_project.go
  touch project_keys.1.8.assign_project_pod.go
  touch project_keys.1.9.unassign_project_pod.go
)

mkdir project_keys.2.lists
(
  # shellcheck disable=SC2164
  cd project_keys.2.lists
  touch project_keys.2.1.add_project_project_list.go
  touch project_keys.2.2.get_project_lists.go
  touch project_keys.2.3.copy_all_pods_in_project_list.go
  touch project_keys.2.4.bulk_copy_pods_in_project_list.go
  touch project_keys.2.5.move_all_pods_in_project_list.go
  touch project_keys.2.6.bulk_move_pods_in_project_list.go
  touch project_keys.2.7.move_project_list.go
  touch project_keys.2.8.get_project_list.go
  touch project_keys.2.9.rename_project_list.go
  touch project_keys.2.10.reorder_project_list.go
  touch project_keys.2.11.archive_project_list.go
)

mkdir registration.1
(
  # shellcheck disable=SC2164
  cd registration.1
  touch registration.1.1.register_new_user_by_email.go
  touch registration.1.2.sign_in_by_email.go
  touch registration.1.3.reset_password.go
  touch registration.1.4.activate_user.go
)

mkdir relations.1
(
  # shellcheck disable=SC2164
  cd relations.1
  touch relations.1.1.get_relations_matching_search_token.go
  touch relations.1.2.get_relations_for_key.go
  touch relations.1.3.get_relations_for_project.go
  touch relations.1.4.get_relations_for_pod.go
  touch relations.1.5.get_relations_for_project_pod.go
  touch relations.1.6.relate_key_to_key.go
  touch relations.1.7.unrelate_key_from_key.go
  touch relations.1.8.relate_project_to_key.go
  touch relations.1.9.unrelate_project_from_key.go
  touch relations.1.10.relate_pod_to_key.go
  touch relations.1.11.unrelate_pod_from_key.go
  touch relations.1.12.relate_pod_to_project.go
  touch relations.1.13.unrelate_pod_from_project.go
  touch relations.1.14.relate_project_to_project.go
  touch relations.1.15.unrelate_project_from_project.go
  touch relations.1.16.relate_pod_to_pod.go
  touch relations.1.17.unrelate_pod_from_pod.go
)

mkdir cd scales.1
(
  # shellcheck disable=SC2164
  cd scales.1
  touch scales.1.1.get_scales.go
  touch scales.1.2.add_scale.go
  touch scales.1.3.get_scale.go
  touch scales.1.4.update_scale.go
  touch scales.1.5.delete_scale.go
  touch scales.1.6.get_projects_using_scale.go
  touch scales.1.7.get_pods_using_scale.go
)

mkdir scheduler.1
(
  # shellcheck disable=SC2164
  cd scheduler.1
  touch scheduler.1.1.get_all_events_in_given_window.go
  touch scheduler.1.2.get_all_events_for_given_day.go
  touch scheduler.1.3.get_standalone_events.go
  touch scheduler.1.4.add_standalone_event.go
  touch scheduler.1.5.update_standalone_event.go
  touch scheduler.1.6.delete_standalone_event.go
)

mkdir search.1
(
  # shellcheck disable=SC2164
  cd search.1
  touch search.1.1.search_key.go_project_or_pod_by_token.go
  touch search.1.2.search_user_by_token.go
)

mkdir teacher_keys.1.students
(
  # shellcheck disable=SC2164
  cd teacher_keys.1.students
  touch teacher_keys.1.1.get_attachment_submissions_as_student.go
  touch teacher_keys.1.2.get_comment_submissions_as_student.go
  touch teacher_keys.1.3.get_all_students_in_a_project.go
  touch teacher_keys.2.Teachers.go
)

mkdir teacher_keys.2.teachers
(
  # shellcheck disable=SC2164
  cd teacher_keys.2.teachers
  touch teacher_keys.2.1.get_student_attachment_submissions_as_teacher.go
  touch teacher_keys.2.2.get_student_comment_submissions_as_teacher.go
  touch teacher_keys.2.3.add_attachment_to_teacher_pod_as_teacher.go
  touch teacher_keys.2.4.add_comment_to_teacher_pod_as_teacher.go
  touch teacher_keys.2.5.get_project_and_pods_grades_for_a_student_as_teacher.go
  touch teacher_keys.2.6.Publish_student_grades_for_a_project.go
  touch teacher_keys.2.7.bulk_publish_pod_grades_for_a_student.go
  touch teacher_keys.2.8.bulk_publish_pod_grades_for_students.go
  touch teacher_keys.2.9.get_project_grades_for_all_students.go
  touch teacher_keys.2.10.get_pod_grades_for_all_students.go
  touch teacher_keys.2.11.assign_grade_to_student.go
  touch teacher_keys.2.12.assign_pod_grade_for_a_student_as_teacher.go
  touch teacher_keys.2.13.get_student_profile.go
)

mkdir templates.1
(
  # shellcheck disable=SC2164
  cd templates.1
  touch templates.1.1.get_key_templates.go
  touch templates.1.2.get_project_templates.go
  touch templates.1.3.get_pod_templates.go
)

mkdir user.1
(
  # shellcheck disable=SC2164
  cd user.1
  touch user.1.1.get_users.go
  touch user.1.2.get_user_by_uuid.go
  touch user.1.3.deactivate_user_account.go
)

mkdir version.1
(
  # shellcheck disable=SC2164
  cd version.1
  touch version.1.1.get_latest_version.go
  touch version.1.2.get_app_status.go
)
