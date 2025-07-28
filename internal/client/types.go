// Code generated from GraphQL schema, DO NOT EDIT.

package client

import (
	"encoding/json"
)


// AuthorIdType represents the AuthorIdType GraphQL type
type AuthorIdType struct {
	Author *authors `json:"author"`
	Errors []string `json:"errors"`
	ID int `json:"id"`
}

// AuthorInputType represents the AuthorInputType GraphQL type
type AuthorInputType struct {
}

// BasicTag represents the BasicTag GraphQL type
type BasicTag struct {
}

// BasicTagType represents the BasicTagType GraphQL type
type BasicTagType struct {
	Category string `json:"category"`
	CategorySlug string `json:"categorySlug"`
	Count int `json:"count"`
	Spoiler bool `json:"spoiler"`
	Tag string `json:"tag"`
	TagSlug string `json:"tagSlug"`
}

// BookDtoInput represents the BookDtoInput GraphQL type
type BookDtoInput struct {
}

// BookDtoType represents the BookDtoType GraphQL type
type BookDtoType struct {
}

// BookIdType represents the BookIdType GraphQL type
type BookIdType struct {
	Book *books `json:"book"`
	Errors []string `json:"errors"`
	ID int `json:"id"`
}

// BookInput represents the BookInput GraphQL type
type BookInput struct {
}

// BookMappingIdType represents the BookMappingIdType GraphQL type
type BookMappingIdType struct {
	Book_mapping *book_mappings `json:"book_mapping"`
	Errors []string `json:"errors"`
	ID int `json:"id"`
}

// BookMappingInput represents the BookMappingInput GraphQL type
type BookMappingInput struct {
}

// BookSeriesDtoInput represents the BookSeriesDtoInput GraphQL type
type BookSeriesDtoInput struct {
}

// Boolean represents the Boolean GraphQL type
type Boolean struct {
}

// Boolean_comparison_exp represents the Boolean_comparison_exp GraphQL type
type Boolean_comparison_exp struct {
}

// CharacterDtoInput represents the CharacterDtoInput GraphQL type
type CharacterDtoInput struct {
}

// CharacterIdType represents the CharacterIdType GraphQL type
type CharacterIdType struct {
	Character *characters `json:"character"`
	Errors []string `json:"errors"`
	ID int `json:"id"`
}

// CharacterInput represents the CharacterInput GraphQL type
type CharacterInput struct {
}

// CollectionImportIdType represents the CollectionImportIdType GraphQL type
type CollectionImportIdType struct {
	Collection_import *collection_imports `json:"collection_import"`
	ID int `json:"id"`
}

// CollectionImportInput represents the CollectionImportInput GraphQL type
type CollectionImportInput struct {
}

// CollectionImportResultIdType represents the CollectionImportResultIdType GraphQL type
type CollectionImportResultIdType struct {
	Collection_import_result *collection_import_results `json:"collection_import_result"`
	ID int `json:"id"`
}

// ContributionInputType represents the ContributionInputType GraphQL type
type ContributionInputType struct {
}

// CreateBookFromPlatformInput represents the CreateBookFromPlatformInput GraphQL type
type CreateBookFromPlatformInput struct {
}

// CreatePromptInput represents the CreatePromptInput GraphQL type
type CreatePromptInput struct {
}

// DatesReadInput represents the DatesReadInput GraphQL type
type DatesReadInput struct {
}

// DeleteFollowedPromptType represents the DeleteFollowedPromptType GraphQL type
type DeleteFollowedPromptType struct {
	Success bool `json:"success"`
}

// DeleteListType represents the DeleteListType GraphQL type
type DeleteListType struct {
	Success bool `json:"success"`
}

// DeleteReadingJournalOutput represents the DeleteReadingJournalOutput GraphQL type
type DeleteReadingJournalOutput struct {
	ID int `json:"id"`
}

// DeleteReadingJournalsOutput represents the DeleteReadingJournalsOutput GraphQL type
type DeleteReadingJournalsOutput struct {
	Ids []int `json:"ids"`
}

// DtoTag represents the DtoTag GraphQL type
type DtoTag struct {
}

// EditionIdType represents the EditionIdType GraphQL type
type EditionIdType struct {
	Edition *editions `json:"edition"`
	Errors []string `json:"errors"`
	ID int `json:"id"`
}

// EditionInput represents the EditionInput GraphQL type
type EditionInput struct {
}

// Float represents the Float GraphQL type
type Float struct {
}

// FollowedListType represents the FollowedListType GraphQL type
type FollowedListType struct {
	Errors []string `json:"errors"`
	Followed_list *followed_lists `json:"followed_list"`
	ID int `json:"id"`
}

// FollowedPromptType represents the FollowedPromptType GraphQL type
type FollowedPromptType struct {
	Errors []string `json:"errors"`
	Followed_prompt *followed_prompts `json:"followed_prompt"`
	ID int `json:"id"`
}

// FollowedUserType represents the FollowedUserType GraphQL type
type FollowedUserType struct {
	Error string `json:"error"`
	Followed_user *users `json:"followed_user"`
	Followed_user_id int `json:"followed_user_id"`
	Followed_users *followed_users `json:"followed_users"`
	ID int `json:"id"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// GoalConditionInput represents the GoalConditionInput GraphQL type
type GoalConditionInput struct {
}

// GoalIdType represents the GoalIdType GraphQL type
type GoalIdType struct {
	Errors []string `json:"errors"`
	Goal *goals `json:"goal"`
	ID int `json:"id"`
}

// GoalInput represents the GoalInput GraphQL type
type GoalInput struct {
}

// ImageIdType represents the ImageIdType GraphQL type
type ImageIdType struct {
	ID int `json:"id"`
	Image *images `json:"image"`
}

// ImageInput represents the ImageInput GraphQL type
type ImageInput struct {
}

// InsertBlockOutput represents the InsertBlockOutput GraphQL type
type InsertBlockOutput struct {
	Error string `json:"error"`
	ID int `json:"id"`
	User_block *user_blocks `json:"user_block"`
}

// Int represents the Int GraphQL type
type Int struct {
}

// Int_comparison_exp represents the Int_comparison_exp GraphQL type
type Int_comparison_exp struct {
}

// LikeDeleteType represents the LikeDeleteType GraphQL type
type LikeDeleteType struct {
	Likes_count int `json:"likes_count"`
}

// LikeType represents the LikeType GraphQL type
type LikeType struct {
	ID int `json:"id"`
	Like *likes `json:"like"`
	Likes_count int `json:"likes_count"`
}

// ListBookDeleteType represents the ListBookDeleteType GraphQL type
type ListBookDeleteType struct {
	ID int `json:"id"`
	List *lists `json:"list"`
	List_id int `json:"list_id"`
}

// ListBookIdType represents the ListBookIdType GraphQL type
type ListBookIdType struct {
	ID int `json:"id"`
	List_book *list_books `json:"list_book"`
}

// ListBookInput represents the ListBookInput GraphQL type
type ListBookInput struct {
}

// ListDeleteType represents the ListDeleteType GraphQL type
type ListDeleteType struct {
	Success bool `json:"success"`
}

// ListIdType represents the ListIdType GraphQL type
type ListIdType struct {
	Errors []string `json:"errors"`
	ID int `json:"id"`
	List *lists `json:"list"`
}

// ListInput represents the ListInput GraphQL type
type ListInput struct {
}

// NewBookIdType represents the NewBookIdType GraphQL type
type NewBookIdType struct {
	Book *books `json:"book"`
	Edition *editions `json:"edition"`
	Edition_id int `json:"edition_id"`
	Errors []string `json:"errors"`
	ID int `json:"id"`
}

// NewsletterStatusType represents the NewsletterStatusType GraphQL type
type NewsletterStatusType struct {
	Subscribed bool `json:"subscribed"`
}

// OptionalEditionIdType represents the OptionalEditionIdType GraphQL type
type OptionalEditionIdType struct {
	Edition *editions `json:"edition"`
	Errors []string `json:"errors"`
	ID int `json:"id"`
}

// PromptAnswerCreateInput represents the PromptAnswerCreateInput GraphQL type
type PromptAnswerCreateInput struct {
}

// PromptAnswerIdType represents the PromptAnswerIdType GraphQL type
type PromptAnswerIdType struct {
	Book_id int `json:"book_id"`
	ID int `json:"id"`
	Prompt_answer *prompt_answers `json:"prompt_answer"`
	Prompt_book *prompt_books_summary `json:"prompt_book"`
	Prompt_id int `json:"prompt_id"`
	User_id int `json:"user_id"`
}

// PromptIdType represents the PromptIdType GraphQL type
type PromptIdType struct {
	Error string `json:"error"`
	ID int `json:"id"`
	Prompt *prompts `json:"prompt"`
}

// PublisherIdType represents the PublisherIdType GraphQL type
type PublisherIdType struct {
	Errors []string `json:"errors"`
	ID int `json:"id"`
	Publisher *publishers `json:"publisher"`
}

// PublisherInputType represents the PublisherInputType GraphQL type
type PublisherInputType struct {
}

// ReadingJournalCreateType represents the ReadingJournalCreateType GraphQL type
type ReadingJournalCreateType struct {
}

// ReadingJournalOutput represents the ReadingJournalOutput GraphQL type
type ReadingJournalOutput struct {
	Errors []string `json:"errors"`
	ID int `json:"id"`
	Reading_journal *reading_journals `json:"reading_journal"`
}

// ReadingJournalUpdateType represents the ReadingJournalUpdateType GraphQL type
type ReadingJournalUpdateType struct {
}

// ReferralType represents the ReferralType GraphQL type
type ReferralType struct {
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Count int `json:"count"`
}

// ReportInput represents the ReportInput GraphQL type
type ReportInput struct {
}

// ReportOutput represents the ReportOutput GraphQL type
type ReportOutput struct {
	Complete bool `json:"complete"`
	Created bool `json:"created"`
	Errors []string `json:"errors"`
}

// SearchOutput represents the SearchOutput GraphQL type
type SearchOutput struct {
	Error string `json:"error"`
	Ids []int `json:"ids"`
	Page int `json:"page"`
	Per_page int `json:"per_page"`
	Query string `json:"query"`
	Query_type string `json:"query_type"`
	Results *json.RawMessage `json:"results"`
}

// SeriesIdType represents the SeriesIdType GraphQL type
type SeriesIdType struct {
	Errors []string `json:"errors"`
	ID int `json:"id"`
	Series *series `json:"series"`
}

// SeriesInput represents the SeriesInput GraphQL type
type SeriesInput struct {
}

// SeriesInputType represents the SeriesInputType GraphQL type
type SeriesInputType struct {
}

// String represents the String GraphQL type
type String struct {
}

// String_comparison_exp represents the String_comparison_exp GraphQL type
type String_comparison_exp struct {
}

// SubscriptionsType represents the SubscriptionsType GraphQL type
type SubscriptionsType struct {
	Billing_portal_url string `json:"billing_portal_url"`
	Membership string `json:"membership"`
	Membership_ends_at *timestamp `json:"membership_ends_at"`
	Monthly_session_id string `json:"monthly_session_id"`
	Monthly_session_url string `json:"monthly_session_url"`
	Payment_system string `json:"payment_system"`
	Yearly_session_id string `json:"yearly_session_id"`
	Yearly_session_url string `json:"yearly_session_url"`
}

// SuccessType represents the SuccessType GraphQL type
type SuccessType struct {
	Success bool `json:"success"`
}

// TagsDtoInput represents the TagsDtoInput GraphQL type
type TagsDtoInput struct {
}

// TagsType represents the TagsType GraphQL type
type TagsType struct {
	Tags []*BasicTagType `json:"tags"`
}

// TrendingBookType represents the TrendingBookType GraphQL type
type TrendingBookType struct {
	Error string `json:"error"`
	Ids []int `json:"ids"`
}

// UpdatePromptInput represents the UpdatePromptInput GraphQL type
type UpdatePromptInput struct {
}

// UserBookCreateInput represents the UserBookCreateInput GraphQL type
type UserBookCreateInput struct {
}

// UserBookDeleteType represents the UserBookDeleteType GraphQL type
type UserBookDeleteType struct {
	Book_id int `json:"book_id"`
	ID int `json:"id"`
	User_book *user_books `json:"user_book"`
	User_id int `json:"user_id"`
}

// UserBookIdType represents the UserBookIdType GraphQL type
type UserBookIdType struct {
	Error string `json:"error"`
	ID int `json:"id"`
	User_book *user_books `json:"user_book"`
}

// UserBookReadIdType represents the UserBookReadIdType GraphQL type
type UserBookReadIdType struct {
	Error string `json:"error"`
	ID int `json:"id"`
	User_book_read *user_book_reads `json:"user_book_read"`
}

// UserBookUpdateInput represents the UserBookUpdateInput GraphQL type
type UserBookUpdateInput struct {
}

// UserBooksReadUpsertType represents the UserBooksReadUpsertType GraphQL type
type UserBooksReadUpsertType struct {
	Error string `json:"error"`
	User_book *user_books `json:"user_book"`
	User_book_id int `json:"user_book_id"`
}

// UserIdType represents the UserIdType GraphQL type
type UserIdType struct {
	Errors []string `json:"errors"`
	ID int `json:"id"`
	User *users `json:"user"`
}

// UserJoinInput represents the UserJoinInput GraphQL type
type UserJoinInput struct {
}

// UserLoginInput represents the UserLoginInput GraphQL type
type UserLoginInput struct {
}

// ValidateReceiptType represents the ValidateReceiptType GraphQL type
type ValidateReceiptType struct {
	Result *json.RawMessage `json:"result"`
	Supporter bool `json:"supporter"`
}

// activities represents the activities GraphQL type
type activities struct {
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Created_at *timestamptz `json:"created_at"`
	Data *json.RawMessage `json:"data"`
	Event string `json:"event"`
	Followers []*followed_users `json:"followers"`
	ID int `json:"id"`
	Likes []*likes `json:"likes"`
	Likes_count int `json:"likes_count"`
	Object_type string `json:"object_type"`
	Original_book_id int `json:"original_book_id"`
	Privacy_setting *privacy_settings `json:"privacy_setting"`
	Privacy_setting_id int `json:"privacy_setting_id"`
	Uid string `json:"uid"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// activities_aggregate_order_by represents the activities_aggregate_order_by GraphQL type
type activities_aggregate_order_by struct {
}

// activities_avg_order_by represents the activities_avg_order_by GraphQL type
type activities_avg_order_by struct {
}

// activities_bool_exp represents the activities_bool_exp GraphQL type
type activities_bool_exp struct {
}

// activities_max_order_by represents the activities_max_order_by GraphQL type
type activities_max_order_by struct {
}

// activities_min_order_by represents the activities_min_order_by GraphQL type
type activities_min_order_by struct {
}

// activities_mutation_response represents the activities_mutation_response GraphQL type
type activities_mutation_response struct {
	Affected_rows int `json:"affected_rows"`
	Returning []*activities `json:"returning"`
}

// activities_order_by represents the activities_order_by GraphQL type
type activities_order_by struct {
}

// activities_select_column represents the activities_select_column GraphQL type
type activities_select_column struct {
}

// activities_stddev_order_by represents the activities_stddev_order_by GraphQL type
type activities_stddev_order_by struct {
}

// activities_stddev_pop_order_by represents the activities_stddev_pop_order_by GraphQL type
type activities_stddev_pop_order_by struct {
}

// activities_stddev_samp_order_by represents the activities_stddev_samp_order_by GraphQL type
type activities_stddev_samp_order_by struct {
}

// activities_stream_cursor_input represents the activities_stream_cursor_input GraphQL type
type activities_stream_cursor_input struct {
}

// activities_stream_cursor_value_input represents the activities_stream_cursor_value_input GraphQL type
type activities_stream_cursor_value_input struct {
}

// activities_sum_order_by represents the activities_sum_order_by GraphQL type
type activities_sum_order_by struct {
}

// activities_var_pop_order_by represents the activities_var_pop_order_by GraphQL type
type activities_var_pop_order_by struct {
}

// activities_var_samp_order_by represents the activities_var_samp_order_by GraphQL type
type activities_var_samp_order_by struct {
}

// activities_variance_order_by represents the activities_variance_order_by GraphQL type
type activities_variance_order_by struct {
}

// activity_feed_args represents the activity_feed_args GraphQL type
type activity_feed_args struct {
}

// activity_foryou_feed_args represents the activity_foryou_feed_args GraphQL type
type activity_foryou_feed_args struct {
}

// authors represents the authors GraphQL type
type authors struct {
	Alias []*authors `json:"alias"`
	Alias_id int `json:"alias_id"`
	Alternate_names *json.RawMessage `json:"alternate_names"`
	Bio string `json:"bio"`
	Books_count int `json:"books_count"`
	Born_date *date `json:"born_date"`
	Born_year int `json:"born_year"`
	Cached_image *json.RawMessage `json:"cached_image"`
	Canonical *authors `json:"canonical"`
	Canonical_id int `json:"canonical_id"`
	Contributions []*contributions `json:"contributions"`
	Contributions_aggregate *contributions_aggregate `json:"contributions_aggregate"`
	Creator *users `json:"creator"`
	Death_date *date `json:"death_date"`
	Death_year int `json:"death_year"`
	Gender_id int `json:"gender_id"`
	ID int `json:"id"`
	Identifiers *json.RawMessage `json:"identifiers"`
	Image *images `json:"image"`
	Image_id int `json:"image_id"`
	Is_bipoc bool `json:"is_bipoc"`
	Is_lgbtq bool `json:"is_lgbtq"`
	Links *json.RawMessage `json:"links"`
	Location string `json:"location"`
	Locked bool `json:"locked"`
	Name string `json:"name"`
	Name_personal string `json:"name_personal"`
	Slug string `json:"slug"`
	State string `json:"state"`
	Title string `json:"title"`
	User_id int `json:"user_id"`
	Users_count int `json:"users_count"`
}

// authors_aggregate_order_by represents the authors_aggregate_order_by GraphQL type
type authors_aggregate_order_by struct {
}

// authors_avg_order_by represents the authors_avg_order_by GraphQL type
type authors_avg_order_by struct {
}

// authors_bool_exp represents the authors_bool_exp GraphQL type
type authors_bool_exp struct {
}

// authors_max_order_by represents the authors_max_order_by GraphQL type
type authors_max_order_by struct {
}

// authors_min_order_by represents the authors_min_order_by GraphQL type
type authors_min_order_by struct {
}

// authors_order_by represents the authors_order_by GraphQL type
type authors_order_by struct {
}

// authors_select_column represents the authors_select_column GraphQL type
type authors_select_column struct {
}

// authors_stddev_order_by represents the authors_stddev_order_by GraphQL type
type authors_stddev_order_by struct {
}

// authors_stddev_pop_order_by represents the authors_stddev_pop_order_by GraphQL type
type authors_stddev_pop_order_by struct {
}

// authors_stddev_samp_order_by represents the authors_stddev_samp_order_by GraphQL type
type authors_stddev_samp_order_by struct {
}

// authors_stream_cursor_input represents the authors_stream_cursor_input GraphQL type
type authors_stream_cursor_input struct {
}

// authors_stream_cursor_value_input represents the authors_stream_cursor_value_input GraphQL type
type authors_stream_cursor_value_input struct {
}

// authors_sum_order_by represents the authors_sum_order_by GraphQL type
type authors_sum_order_by struct {
}

// authors_var_pop_order_by represents the authors_var_pop_order_by GraphQL type
type authors_var_pop_order_by struct {
}

// authors_var_samp_order_by represents the authors_var_samp_order_by GraphQL type
type authors_var_samp_order_by struct {
}

// authors_variance_order_by represents the authors_variance_order_by GraphQL type
type authors_variance_order_by struct {
}

// bigint represents the bigint GraphQL type
type bigint struct {
}

// bigint_comparison_exp represents the bigint_comparison_exp GraphQL type
type bigint_comparison_exp struct {
}

// book_categories represents the book_categories GraphQL type
type book_categories struct {
	ID *bigint `json:"id"`
	Name string `json:"name"`
}

// book_categories_bool_exp represents the book_categories_bool_exp GraphQL type
type book_categories_bool_exp struct {
}

// book_categories_order_by represents the book_categories_order_by GraphQL type
type book_categories_order_by struct {
}

// book_categories_select_column represents the book_categories_select_column GraphQL type
type book_categories_select_column struct {
}

// book_categories_stream_cursor_input represents the book_categories_stream_cursor_input GraphQL type
type book_categories_stream_cursor_input struct {
}

// book_categories_stream_cursor_value_input represents the book_categories_stream_cursor_value_input GraphQL type
type book_categories_stream_cursor_value_input struct {
}

// book_characters represents the book_characters GraphQL type
type book_characters struct {
	Book *books `json:"book"`
	Book_id *bigint `json:"book_id"`
	Character *characters `json:"character"`
	Character_id *bigint `json:"character_id"`
	ID *bigint `json:"id"`
	Only_mentioned bool `json:"only_mentioned"`
	Position int `json:"position"`
	Spoiler bool `json:"spoiler"`
}

// book_characters_aggregate_order_by represents the book_characters_aggregate_order_by GraphQL type
type book_characters_aggregate_order_by struct {
}

// book_characters_avg_order_by represents the book_characters_avg_order_by GraphQL type
type book_characters_avg_order_by struct {
}

// book_characters_bool_exp represents the book_characters_bool_exp GraphQL type
type book_characters_bool_exp struct {
}

// book_characters_max_order_by represents the book_characters_max_order_by GraphQL type
type book_characters_max_order_by struct {
}

// book_characters_min_order_by represents the book_characters_min_order_by GraphQL type
type book_characters_min_order_by struct {
}

// book_characters_order_by represents the book_characters_order_by GraphQL type
type book_characters_order_by struct {
}

// book_characters_select_column represents the book_characters_select_column GraphQL type
type book_characters_select_column struct {
}

// book_characters_stddev_order_by represents the book_characters_stddev_order_by GraphQL type
type book_characters_stddev_order_by struct {
}

// book_characters_stddev_pop_order_by represents the book_characters_stddev_pop_order_by GraphQL type
type book_characters_stddev_pop_order_by struct {
}

// book_characters_stddev_samp_order_by represents the book_characters_stddev_samp_order_by GraphQL type
type book_characters_stddev_samp_order_by struct {
}

// book_characters_stream_cursor_input represents the book_characters_stream_cursor_input GraphQL type
type book_characters_stream_cursor_input struct {
}

// book_characters_stream_cursor_value_input represents the book_characters_stream_cursor_value_input GraphQL type
type book_characters_stream_cursor_value_input struct {
}

// book_characters_sum_order_by represents the book_characters_sum_order_by GraphQL type
type book_characters_sum_order_by struct {
}

// book_characters_var_pop_order_by represents the book_characters_var_pop_order_by GraphQL type
type book_characters_var_pop_order_by struct {
}

// book_characters_var_samp_order_by represents the book_characters_var_samp_order_by GraphQL type
type book_characters_var_samp_order_by struct {
}

// book_characters_variance_order_by represents the book_characters_variance_order_by GraphQL type
type book_characters_variance_order_by struct {
}

// book_collections represents the book_collections GraphQL type
type book_collections struct {
	Book_id int `json:"book_id"`
	Child_book_id int `json:"child_book_id"`
	ID *bigint `json:"id"`
	Position int `json:"position"`
}

// book_collections_bool_exp represents the book_collections_bool_exp GraphQL type
type book_collections_bool_exp struct {
}

// book_collections_order_by represents the book_collections_order_by GraphQL type
type book_collections_order_by struct {
}

// book_collections_select_column represents the book_collections_select_column GraphQL type
type book_collections_select_column struct {
}

// book_collections_stream_cursor_input represents the book_collections_stream_cursor_input GraphQL type
type book_collections_stream_cursor_input struct {
}

// book_collections_stream_cursor_value_input represents the book_collections_stream_cursor_value_input GraphQL type
type book_collections_stream_cursor_value_input struct {
}

// book_mappings represents the book_mappings GraphQL type
type book_mappings struct {
	Attempts int `json:"attempts"`
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Created_at *timestamptz `json:"created_at"`
	Dto_external *json.RawMessage `json:"dto_external"`
	Edition *editions `json:"edition"`
	Edition_id int `json:"edition_id"`
	External_data_id int `json:"external_data_id"`
	External_id string `json:"external_id"`
	ID int `json:"id"`
	Loaded bool `json:"loaded"`
	Loaded_at *timestamp `json:"loaded_at"`
	Normalized_at *timestamp `json:"normalized_at"`
	Original_book_id int `json:"original_book_id"`
	Platform *platforms `json:"platform"`
	Platform_id int `json:"platform_id"`
	State string `json:"state"`
	Updated_at *timestamptz `json:"updated_at"`
	Verified bool `json:"verified"`
	Verified_at *timestamp `json:"verified_at"`
}

// book_mappings_aggregate_order_by represents the book_mappings_aggregate_order_by GraphQL type
type book_mappings_aggregate_order_by struct {
}

// book_mappings_avg_order_by represents the book_mappings_avg_order_by GraphQL type
type book_mappings_avg_order_by struct {
}

// book_mappings_bool_exp represents the book_mappings_bool_exp GraphQL type
type book_mappings_bool_exp struct {
}

// book_mappings_max_order_by represents the book_mappings_max_order_by GraphQL type
type book_mappings_max_order_by struct {
}

// book_mappings_min_order_by represents the book_mappings_min_order_by GraphQL type
type book_mappings_min_order_by struct {
}

// book_mappings_order_by represents the book_mappings_order_by GraphQL type
type book_mappings_order_by struct {
}

// book_mappings_select_column represents the book_mappings_select_column GraphQL type
type book_mappings_select_column struct {
}

// book_mappings_stddev_order_by represents the book_mappings_stddev_order_by GraphQL type
type book_mappings_stddev_order_by struct {
}

// book_mappings_stddev_pop_order_by represents the book_mappings_stddev_pop_order_by GraphQL type
type book_mappings_stddev_pop_order_by struct {
}

// book_mappings_stddev_samp_order_by represents the book_mappings_stddev_samp_order_by GraphQL type
type book_mappings_stddev_samp_order_by struct {
}

// book_mappings_stream_cursor_input represents the book_mappings_stream_cursor_input GraphQL type
type book_mappings_stream_cursor_input struct {
}

// book_mappings_stream_cursor_value_input represents the book_mappings_stream_cursor_value_input GraphQL type
type book_mappings_stream_cursor_value_input struct {
}

// book_mappings_sum_order_by represents the book_mappings_sum_order_by GraphQL type
type book_mappings_sum_order_by struct {
}

// book_mappings_var_pop_order_by represents the book_mappings_var_pop_order_by GraphQL type
type book_mappings_var_pop_order_by struct {
}

// book_mappings_var_samp_order_by represents the book_mappings_var_samp_order_by GraphQL type
type book_mappings_var_samp_order_by struct {
}

// book_mappings_variance_order_by represents the book_mappings_variance_order_by GraphQL type
type book_mappings_variance_order_by struct {
}

// book_series represents the book_series GraphQL type
type book_series struct {
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Created_at *timestamp `json:"created_at"`
	Details string `json:"details"`
	Featured bool `json:"featured"`
	ID *bigint `json:"id"`
	Position *float8 `json:"position"`
	Series *series `json:"series"`
	Series_id int `json:"series_id"`
	Updated_at *timestamp `json:"updated_at"`
}

// book_series_aggregate represents the book_series_aggregate GraphQL type
type book_series_aggregate struct {
	Aggregate *book_series_aggregate_fields `json:"aggregate"`
	Nodes []*book_series `json:"nodes"`
}

// book_series_aggregate_bool_exp represents the book_series_aggregate_bool_exp GraphQL type
type book_series_aggregate_bool_exp struct {
}

// book_series_aggregate_bool_exp_avg represents the book_series_aggregate_bool_exp_avg GraphQL type
type book_series_aggregate_bool_exp_avg struct {
}

// book_series_aggregate_bool_exp_bool_and represents the book_series_aggregate_bool_exp_bool_and GraphQL type
type book_series_aggregate_bool_exp_bool_and struct {
}

// book_series_aggregate_bool_exp_bool_or represents the book_series_aggregate_bool_exp_bool_or GraphQL type
type book_series_aggregate_bool_exp_bool_or struct {
}

// book_series_aggregate_bool_exp_corr represents the book_series_aggregate_bool_exp_corr GraphQL type
type book_series_aggregate_bool_exp_corr struct {
}

// book_series_aggregate_bool_exp_corr_arguments represents the book_series_aggregate_bool_exp_corr_arguments GraphQL type
type book_series_aggregate_bool_exp_corr_arguments struct {
}

// book_series_aggregate_bool_exp_count represents the book_series_aggregate_bool_exp_count GraphQL type
type book_series_aggregate_bool_exp_count struct {
}

// book_series_aggregate_bool_exp_covar_samp represents the book_series_aggregate_bool_exp_covar_samp GraphQL type
type book_series_aggregate_bool_exp_covar_samp struct {
}

// book_series_aggregate_bool_exp_covar_samp_arguments represents the book_series_aggregate_bool_exp_covar_samp_arguments GraphQL type
type book_series_aggregate_bool_exp_covar_samp_arguments struct {
}

// book_series_aggregate_bool_exp_max represents the book_series_aggregate_bool_exp_max GraphQL type
type book_series_aggregate_bool_exp_max struct {
}

// book_series_aggregate_bool_exp_min represents the book_series_aggregate_bool_exp_min GraphQL type
type book_series_aggregate_bool_exp_min struct {
}

// book_series_aggregate_bool_exp_stddev_samp represents the book_series_aggregate_bool_exp_stddev_samp GraphQL type
type book_series_aggregate_bool_exp_stddev_samp struct {
}

// book_series_aggregate_bool_exp_sum represents the book_series_aggregate_bool_exp_sum GraphQL type
type book_series_aggregate_bool_exp_sum struct {
}

// book_series_aggregate_bool_exp_var_samp represents the book_series_aggregate_bool_exp_var_samp GraphQL type
type book_series_aggregate_bool_exp_var_samp struct {
}

// book_series_aggregate_fields represents the book_series_aggregate_fields GraphQL type
type book_series_aggregate_fields struct {
	Avg *book_series_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *book_series_max_fields `json:"max"`
	Min *book_series_min_fields `json:"min"`
	Stddev *book_series_stddev_fields `json:"stddev"`
	Stddev_pop *book_series_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *book_series_stddev_samp_fields `json:"stddev_samp"`
	Sum *book_series_sum_fields `json:"sum"`
	Var_pop *book_series_var_pop_fields `json:"var_pop"`
	Var_samp *book_series_var_samp_fields `json:"var_samp"`
	Variance *book_series_variance_fields `json:"variance"`
}

// book_series_aggregate_order_by represents the book_series_aggregate_order_by GraphQL type
type book_series_aggregate_order_by struct {
}

// book_series_avg_fields represents the book_series_avg_fields GraphQL type
type book_series_avg_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Position float64 `json:"position"`
	Series_id float64 `json:"series_id"`
}

// book_series_avg_order_by represents the book_series_avg_order_by GraphQL type
type book_series_avg_order_by struct {
}

// book_series_bool_exp represents the book_series_bool_exp GraphQL type
type book_series_bool_exp struct {
}

// book_series_max_fields represents the book_series_max_fields GraphQL type
type book_series_max_fields struct {
	Book_id int `json:"book_id"`
	Created_at *timestamp `json:"created_at"`
	Details string `json:"details"`
	ID *bigint `json:"id"`
	Position *float8 `json:"position"`
	Series_id int `json:"series_id"`
	Updated_at *timestamp `json:"updated_at"`
}

// book_series_max_order_by represents the book_series_max_order_by GraphQL type
type book_series_max_order_by struct {
}

// book_series_min_fields represents the book_series_min_fields GraphQL type
type book_series_min_fields struct {
	Book_id int `json:"book_id"`
	Created_at *timestamp `json:"created_at"`
	Details string `json:"details"`
	ID *bigint `json:"id"`
	Position *float8 `json:"position"`
	Series_id int `json:"series_id"`
	Updated_at *timestamp `json:"updated_at"`
}

// book_series_min_order_by represents the book_series_min_order_by GraphQL type
type book_series_min_order_by struct {
}

// book_series_order_by represents the book_series_order_by GraphQL type
type book_series_order_by struct {
}

// book_series_select_column represents the book_series_select_column GraphQL type
type book_series_select_column struct {
}

// book_series_select_column_book_series_aggregate_bool_exp_avg_arguments_columns represents the book_series_select_column_book_series_aggregate_bool_exp_avg_arguments_columns GraphQL type
type book_series_select_column_book_series_aggregate_bool_exp_avg_arguments_columns struct {
}

// book_series_select_column_book_series_aggregate_bool_exp_bool_and_arguments_columns represents the book_series_select_column_book_series_aggregate_bool_exp_bool_and_arguments_columns GraphQL type
type book_series_select_column_book_series_aggregate_bool_exp_bool_and_arguments_columns struct {
}

// book_series_select_column_book_series_aggregate_bool_exp_bool_or_arguments_columns represents the book_series_select_column_book_series_aggregate_bool_exp_bool_or_arguments_columns GraphQL type
type book_series_select_column_book_series_aggregate_bool_exp_bool_or_arguments_columns struct {
}

// book_series_select_column_book_series_aggregate_bool_exp_corr_arguments_columns represents the book_series_select_column_book_series_aggregate_bool_exp_corr_arguments_columns GraphQL type
type book_series_select_column_book_series_aggregate_bool_exp_corr_arguments_columns struct {
}

// book_series_select_column_book_series_aggregate_bool_exp_covar_samp_arguments_columns represents the book_series_select_column_book_series_aggregate_bool_exp_covar_samp_arguments_columns GraphQL type
type book_series_select_column_book_series_aggregate_bool_exp_covar_samp_arguments_columns struct {
}

// book_series_select_column_book_series_aggregate_bool_exp_max_arguments_columns represents the book_series_select_column_book_series_aggregate_bool_exp_max_arguments_columns GraphQL type
type book_series_select_column_book_series_aggregate_bool_exp_max_arguments_columns struct {
}

// book_series_select_column_book_series_aggregate_bool_exp_min_arguments_columns represents the book_series_select_column_book_series_aggregate_bool_exp_min_arguments_columns GraphQL type
type book_series_select_column_book_series_aggregate_bool_exp_min_arguments_columns struct {
}

// book_series_select_column_book_series_aggregate_bool_exp_stddev_samp_arguments_columns represents the book_series_select_column_book_series_aggregate_bool_exp_stddev_samp_arguments_columns GraphQL type
type book_series_select_column_book_series_aggregate_bool_exp_stddev_samp_arguments_columns struct {
}

// book_series_select_column_book_series_aggregate_bool_exp_sum_arguments_columns represents the book_series_select_column_book_series_aggregate_bool_exp_sum_arguments_columns GraphQL type
type book_series_select_column_book_series_aggregate_bool_exp_sum_arguments_columns struct {
}

// book_series_select_column_book_series_aggregate_bool_exp_var_samp_arguments_columns represents the book_series_select_column_book_series_aggregate_bool_exp_var_samp_arguments_columns GraphQL type
type book_series_select_column_book_series_aggregate_bool_exp_var_samp_arguments_columns struct {
}

// book_series_stddev_fields represents the book_series_stddev_fields GraphQL type
type book_series_stddev_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Position float64 `json:"position"`
	Series_id float64 `json:"series_id"`
}

// book_series_stddev_order_by represents the book_series_stddev_order_by GraphQL type
type book_series_stddev_order_by struct {
}

// book_series_stddev_pop_fields represents the book_series_stddev_pop_fields GraphQL type
type book_series_stddev_pop_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Position float64 `json:"position"`
	Series_id float64 `json:"series_id"`
}

// book_series_stddev_pop_order_by represents the book_series_stddev_pop_order_by GraphQL type
type book_series_stddev_pop_order_by struct {
}

// book_series_stddev_samp_fields represents the book_series_stddev_samp_fields GraphQL type
type book_series_stddev_samp_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Position float64 `json:"position"`
	Series_id float64 `json:"series_id"`
}

// book_series_stddev_samp_order_by represents the book_series_stddev_samp_order_by GraphQL type
type book_series_stddev_samp_order_by struct {
}

// book_series_stream_cursor_input represents the book_series_stream_cursor_input GraphQL type
type book_series_stream_cursor_input struct {
}

// book_series_stream_cursor_value_input represents the book_series_stream_cursor_value_input GraphQL type
type book_series_stream_cursor_value_input struct {
}

// book_series_sum_fields represents the book_series_sum_fields GraphQL type
type book_series_sum_fields struct {
	Book_id int `json:"book_id"`
	ID *bigint `json:"id"`
	Position *float8 `json:"position"`
	Series_id int `json:"series_id"`
}

// book_series_sum_order_by represents the book_series_sum_order_by GraphQL type
type book_series_sum_order_by struct {
}

// book_series_var_pop_fields represents the book_series_var_pop_fields GraphQL type
type book_series_var_pop_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Position float64 `json:"position"`
	Series_id float64 `json:"series_id"`
}

// book_series_var_pop_order_by represents the book_series_var_pop_order_by GraphQL type
type book_series_var_pop_order_by struct {
}

// book_series_var_samp_fields represents the book_series_var_samp_fields GraphQL type
type book_series_var_samp_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Position float64 `json:"position"`
	Series_id float64 `json:"series_id"`
}

// book_series_var_samp_order_by represents the book_series_var_samp_order_by GraphQL type
type book_series_var_samp_order_by struct {
}

// book_series_variance_fields represents the book_series_variance_fields GraphQL type
type book_series_variance_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Position float64 `json:"position"`
	Series_id float64 `json:"series_id"`
}

// book_series_variance_order_by represents the book_series_variance_order_by GraphQL type
type book_series_variance_order_by struct {
}

// book_statuses represents the book_statuses GraphQL type
type book_statuses struct {
	Books []*books `json:"books"`
	Books_aggregate *books_aggregate `json:"books_aggregate"`
	ID *smallint `json:"id"`
	Name string `json:"name"`
}

// book_statuses_bool_exp represents the book_statuses_bool_exp GraphQL type
type book_statuses_bool_exp struct {
}

// book_statuses_order_by represents the book_statuses_order_by GraphQL type
type book_statuses_order_by struct {
}

// book_statuses_select_column represents the book_statuses_select_column GraphQL type
type book_statuses_select_column struct {
}

// book_statuses_stream_cursor_input represents the book_statuses_stream_cursor_input GraphQL type
type book_statuses_stream_cursor_input struct {
}

// book_statuses_stream_cursor_value_input represents the book_statuses_stream_cursor_value_input GraphQL type
type book_statuses_stream_cursor_value_input struct {
}

// bookles represents the bookles GraphQL type
type bookles struct {
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Created_at *timestamp `json:"created_at"`
	Date *date `json:"date"`
	ID *bigint `json:"id"`
}

// bookles_bool_exp represents the bookles_bool_exp GraphQL type
type bookles_bool_exp struct {
}

// bookles_order_by represents the bookles_order_by GraphQL type
type bookles_order_by struct {
}

// bookles_select_column represents the bookles_select_column GraphQL type
type bookles_select_column struct {
}

// bookles_stream_cursor_input represents the bookles_stream_cursor_input GraphQL type
type bookles_stream_cursor_input struct {
}

// bookles_stream_cursor_value_input represents the bookles_stream_cursor_value_input GraphQL type
type bookles_stream_cursor_value_input struct {
}

// books represents the books GraphQL type
type books struct {
	Activities_count int `json:"activities_count"`
	Alternative_titles *json.RawMessage `json:"alternative_titles"`
	Audio_seconds int `json:"audio_seconds"`
	Book_category_id int `json:"book_category_id"`
	Book_characters []*book_characters `json:"book_characters"`
	Book_mappings []*book_mappings `json:"book_mappings"`
	Book_series []*book_series `json:"book_series"`
	Book_series_aggregate *book_series_aggregate `json:"book_series_aggregate"`
	Book_status *book_statuses `json:"book_status"`
	Book_status_id *smallint `json:"book_status_id"`
	Cached_contributors *json.RawMessage `json:"cached_contributors"`
	Cached_featured_series *json.RawMessage `json:"cached_featured_series"`
	Cached_header_image *json.RawMessage `json:"cached_header_image"`
	Cached_image *json.RawMessage `json:"cached_image"`
	Cached_tags *json.RawMessage `json:"cached_tags"`
	Canonical *books `json:"canonical"`
	Canonical_id int `json:"canonical_id"`
	Collection_import_results []*collection_import_results `json:"collection_import_results"`
	Compilation bool `json:"compilation"`
	Contributions []*contributions `json:"contributions"`
	Contributions_aggregate *contributions_aggregate `json:"contributions_aggregate"`
	Created_at *timestamp `json:"created_at"`
	Created_by_user_id int `json:"created_by_user_id"`
	Default_audio_edition *editions `json:"default_audio_edition"`
	Default_audio_edition_id int `json:"default_audio_edition_id"`
	Default_cover_edition *editions `json:"default_cover_edition"`
	Default_cover_edition_id int `json:"default_cover_edition_id"`
	Default_ebook_edition *editions `json:"default_ebook_edition"`
	Default_ebook_edition_id int `json:"default_ebook_edition_id"`
	Default_physical_edition *editions `json:"default_physical_edition"`
	Default_physical_edition_id int `json:"default_physical_edition_id"`
	Description string `json:"description"`
	Dto *json.RawMessage `json:"dto"`
	Dto_combined *json.RawMessage `json:"dto_combined"`
	Dto_external *json.RawMessage `json:"dto_external"`
	Editions []*editions `json:"editions"`
	Editions_count int `json:"editions_count"`
	Featured_book_series *book_series `json:"featured_book_series"`
	Featured_book_series_id int `json:"featured_book_series_id"`
	Header_image_id int `json:"header_image_id"`
	Headline string `json:"headline"`
	ID int `json:"id"`
	Image *images `json:"image"`
	Image_id int `json:"image_id"`
	Images []*images `json:"images"`
	Import_platform_id int `json:"import_platform_id"`
	Journals_count int `json:"journals_count"`
	Links *json.RawMessage `json:"links"`
	List_books []*list_books `json:"list_books"`
	List_books_aggregate *list_books_aggregate `json:"list_books_aggregate"`
	Lists_count int `json:"lists_count"`
	Literary_type_id int `json:"literary_type_id"`
	Locked bool `json:"locked"`
	Pages int `json:"pages"`
	Prompt_answers []*prompt_answers `json:"prompt_answers"`
	Prompt_answers_aggregate *prompt_answers_aggregate `json:"prompt_answers_aggregate"`
	Prompt_summaries []*prompt_books_summary `json:"prompt_summaries"`
	Prompts_count int `json:"prompts_count"`
	Rating *numeric `json:"rating"`
	Ratings_count int `json:"ratings_count"`
	Ratings_distribution *json.RawMessage `json:"ratings_distribution"`
	Recommendations []*recommendations `json:"recommendations"`
	Release_date *date `json:"release_date"`
	Release_year int `json:"release_year"`
	Reviews_count int `json:"reviews_count"`
	Slug string `json:"slug"`
	State string `json:"state"`
	Subtitle string `json:"subtitle"`
	Taggable_counts []*taggable_counts `json:"taggable_counts"`
	Taggings []*taggings `json:"taggings"`
	Taggings_aggregate *taggings_aggregate `json:"taggings_aggregate"`
	Title string `json:"title"`
	Updated_at *timestamptz `json:"updated_at"`
	User_added bool `json:"user_added"`
	User_books []*user_books `json:"user_books"`
	User_books_aggregate *user_books_aggregate `json:"user_books_aggregate"`
	Users_count int `json:"users_count"`
	Users_read_count int `json:"users_read_count"`
}

// books_aggregate represents the books_aggregate GraphQL type
type books_aggregate struct {
	Aggregate *books_aggregate_fields `json:"aggregate"`
	Nodes []*books `json:"nodes"`
}

// books_aggregate_bool_exp represents the books_aggregate_bool_exp GraphQL type
type books_aggregate_bool_exp struct {
}

// books_aggregate_bool_exp_bool_and represents the books_aggregate_bool_exp_bool_and GraphQL type
type books_aggregate_bool_exp_bool_and struct {
}

// books_aggregate_bool_exp_bool_or represents the books_aggregate_bool_exp_bool_or GraphQL type
type books_aggregate_bool_exp_bool_or struct {
}

// books_aggregate_bool_exp_count represents the books_aggregate_bool_exp_count GraphQL type
type books_aggregate_bool_exp_count struct {
}

// books_aggregate_fields represents the books_aggregate_fields GraphQL type
type books_aggregate_fields struct {
	Avg *books_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *books_max_fields `json:"max"`
	Min *books_min_fields `json:"min"`
	Stddev *books_stddev_fields `json:"stddev"`
	Stddev_pop *books_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *books_stddev_samp_fields `json:"stddev_samp"`
	Sum *books_sum_fields `json:"sum"`
	Var_pop *books_var_pop_fields `json:"var_pop"`
	Var_samp *books_var_samp_fields `json:"var_samp"`
	Variance *books_variance_fields `json:"variance"`
}

// books_aggregate_order_by represents the books_aggregate_order_by GraphQL type
type books_aggregate_order_by struct {
}

// books_avg_fields represents the books_avg_fields GraphQL type
type books_avg_fields struct {
	Activities_count float64 `json:"activities_count"`
	Audio_seconds float64 `json:"audio_seconds"`
	Book_category_id float64 `json:"book_category_id"`
	Book_status_id float64 `json:"book_status_id"`
	Canonical_id float64 `json:"canonical_id"`
	Created_by_user_id float64 `json:"created_by_user_id"`
	Default_audio_edition_id float64 `json:"default_audio_edition_id"`
	Default_cover_edition_id float64 `json:"default_cover_edition_id"`
	Default_ebook_edition_id float64 `json:"default_ebook_edition_id"`
	Default_physical_edition_id float64 `json:"default_physical_edition_id"`
	Editions_count float64 `json:"editions_count"`
	Featured_book_series_id float64 `json:"featured_book_series_id"`
	Header_image_id float64 `json:"header_image_id"`
	ID float64 `json:"id"`
	Image_id float64 `json:"image_id"`
	Import_platform_id float64 `json:"import_platform_id"`
	Journals_count float64 `json:"journals_count"`
	Lists_count float64 `json:"lists_count"`
	Literary_type_id float64 `json:"literary_type_id"`
	Pages float64 `json:"pages"`
	Prompts_count float64 `json:"prompts_count"`
	Rating float64 `json:"rating"`
	Ratings_count float64 `json:"ratings_count"`
	Release_year float64 `json:"release_year"`
	Reviews_count float64 `json:"reviews_count"`
	Users_count float64 `json:"users_count"`
	Users_read_count float64 `json:"users_read_count"`
}

// books_avg_order_by represents the books_avg_order_by GraphQL type
type books_avg_order_by struct {
}

// books_bool_exp represents the books_bool_exp GraphQL type
type books_bool_exp struct {
}

// books_max_fields represents the books_max_fields GraphQL type
type books_max_fields struct {
	Activities_count int `json:"activities_count"`
	Audio_seconds int `json:"audio_seconds"`
	Book_category_id int `json:"book_category_id"`
	Book_status_id *smallint `json:"book_status_id"`
	Canonical_id int `json:"canonical_id"`
	Created_at *timestamp `json:"created_at"`
	Created_by_user_id int `json:"created_by_user_id"`
	Default_audio_edition_id int `json:"default_audio_edition_id"`
	Default_cover_edition_id int `json:"default_cover_edition_id"`
	Default_ebook_edition_id int `json:"default_ebook_edition_id"`
	Default_physical_edition_id int `json:"default_physical_edition_id"`
	Description string `json:"description"`
	Editions_count int `json:"editions_count"`
	Featured_book_series_id int `json:"featured_book_series_id"`
	Header_image_id int `json:"header_image_id"`
	Headline string `json:"headline"`
	ID int `json:"id"`
	Image_id int `json:"image_id"`
	Import_platform_id int `json:"import_platform_id"`
	Journals_count int `json:"journals_count"`
	Lists_count int `json:"lists_count"`
	Literary_type_id int `json:"literary_type_id"`
	Pages int `json:"pages"`
	Prompts_count int `json:"prompts_count"`
	Rating *numeric `json:"rating"`
	Ratings_count int `json:"ratings_count"`
	Release_date *date `json:"release_date"`
	Release_year int `json:"release_year"`
	Reviews_count int `json:"reviews_count"`
	Slug string `json:"slug"`
	State string `json:"state"`
	Subtitle string `json:"subtitle"`
	Title string `json:"title"`
	Updated_at *timestamptz `json:"updated_at"`
	Users_count int `json:"users_count"`
	Users_read_count int `json:"users_read_count"`
}

// books_max_order_by represents the books_max_order_by GraphQL type
type books_max_order_by struct {
}

// books_min_fields represents the books_min_fields GraphQL type
type books_min_fields struct {
	Activities_count int `json:"activities_count"`
	Audio_seconds int `json:"audio_seconds"`
	Book_category_id int `json:"book_category_id"`
	Book_status_id *smallint `json:"book_status_id"`
	Canonical_id int `json:"canonical_id"`
	Created_at *timestamp `json:"created_at"`
	Created_by_user_id int `json:"created_by_user_id"`
	Default_audio_edition_id int `json:"default_audio_edition_id"`
	Default_cover_edition_id int `json:"default_cover_edition_id"`
	Default_ebook_edition_id int `json:"default_ebook_edition_id"`
	Default_physical_edition_id int `json:"default_physical_edition_id"`
	Description string `json:"description"`
	Editions_count int `json:"editions_count"`
	Featured_book_series_id int `json:"featured_book_series_id"`
	Header_image_id int `json:"header_image_id"`
	Headline string `json:"headline"`
	ID int `json:"id"`
	Image_id int `json:"image_id"`
	Import_platform_id int `json:"import_platform_id"`
	Journals_count int `json:"journals_count"`
	Lists_count int `json:"lists_count"`
	Literary_type_id int `json:"literary_type_id"`
	Pages int `json:"pages"`
	Prompts_count int `json:"prompts_count"`
	Rating *numeric `json:"rating"`
	Ratings_count int `json:"ratings_count"`
	Release_date *date `json:"release_date"`
	Release_year int `json:"release_year"`
	Reviews_count int `json:"reviews_count"`
	Slug string `json:"slug"`
	State string `json:"state"`
	Subtitle string `json:"subtitle"`
	Title string `json:"title"`
	Updated_at *timestamptz `json:"updated_at"`
	Users_count int `json:"users_count"`
	Users_read_count int `json:"users_read_count"`
}

// books_min_order_by represents the books_min_order_by GraphQL type
type books_min_order_by struct {
}

// books_order_by represents the books_order_by GraphQL type
type books_order_by struct {
}

// books_select_column represents the books_select_column GraphQL type
type books_select_column struct {
}

// books_select_column_books_aggregate_bool_exp_bool_and_arguments_columns represents the books_select_column_books_aggregate_bool_exp_bool_and_arguments_columns GraphQL type
type books_select_column_books_aggregate_bool_exp_bool_and_arguments_columns struct {
}

// books_select_column_books_aggregate_bool_exp_bool_or_arguments_columns represents the books_select_column_books_aggregate_bool_exp_bool_or_arguments_columns GraphQL type
type books_select_column_books_aggregate_bool_exp_bool_or_arguments_columns struct {
}

// books_stddev_fields represents the books_stddev_fields GraphQL type
type books_stddev_fields struct {
	Activities_count float64 `json:"activities_count"`
	Audio_seconds float64 `json:"audio_seconds"`
	Book_category_id float64 `json:"book_category_id"`
	Book_status_id float64 `json:"book_status_id"`
	Canonical_id float64 `json:"canonical_id"`
	Created_by_user_id float64 `json:"created_by_user_id"`
	Default_audio_edition_id float64 `json:"default_audio_edition_id"`
	Default_cover_edition_id float64 `json:"default_cover_edition_id"`
	Default_ebook_edition_id float64 `json:"default_ebook_edition_id"`
	Default_physical_edition_id float64 `json:"default_physical_edition_id"`
	Editions_count float64 `json:"editions_count"`
	Featured_book_series_id float64 `json:"featured_book_series_id"`
	Header_image_id float64 `json:"header_image_id"`
	ID float64 `json:"id"`
	Image_id float64 `json:"image_id"`
	Import_platform_id float64 `json:"import_platform_id"`
	Journals_count float64 `json:"journals_count"`
	Lists_count float64 `json:"lists_count"`
	Literary_type_id float64 `json:"literary_type_id"`
	Pages float64 `json:"pages"`
	Prompts_count float64 `json:"prompts_count"`
	Rating float64 `json:"rating"`
	Ratings_count float64 `json:"ratings_count"`
	Release_year float64 `json:"release_year"`
	Reviews_count float64 `json:"reviews_count"`
	Users_count float64 `json:"users_count"`
	Users_read_count float64 `json:"users_read_count"`
}

// books_stddev_order_by represents the books_stddev_order_by GraphQL type
type books_stddev_order_by struct {
}

// books_stddev_pop_fields represents the books_stddev_pop_fields GraphQL type
type books_stddev_pop_fields struct {
	Activities_count float64 `json:"activities_count"`
	Audio_seconds float64 `json:"audio_seconds"`
	Book_category_id float64 `json:"book_category_id"`
	Book_status_id float64 `json:"book_status_id"`
	Canonical_id float64 `json:"canonical_id"`
	Created_by_user_id float64 `json:"created_by_user_id"`
	Default_audio_edition_id float64 `json:"default_audio_edition_id"`
	Default_cover_edition_id float64 `json:"default_cover_edition_id"`
	Default_ebook_edition_id float64 `json:"default_ebook_edition_id"`
	Default_physical_edition_id float64 `json:"default_physical_edition_id"`
	Editions_count float64 `json:"editions_count"`
	Featured_book_series_id float64 `json:"featured_book_series_id"`
	Header_image_id float64 `json:"header_image_id"`
	ID float64 `json:"id"`
	Image_id float64 `json:"image_id"`
	Import_platform_id float64 `json:"import_platform_id"`
	Journals_count float64 `json:"journals_count"`
	Lists_count float64 `json:"lists_count"`
	Literary_type_id float64 `json:"literary_type_id"`
	Pages float64 `json:"pages"`
	Prompts_count float64 `json:"prompts_count"`
	Rating float64 `json:"rating"`
	Ratings_count float64 `json:"ratings_count"`
	Release_year float64 `json:"release_year"`
	Reviews_count float64 `json:"reviews_count"`
	Users_count float64 `json:"users_count"`
	Users_read_count float64 `json:"users_read_count"`
}

// books_stddev_pop_order_by represents the books_stddev_pop_order_by GraphQL type
type books_stddev_pop_order_by struct {
}

// books_stddev_samp_fields represents the books_stddev_samp_fields GraphQL type
type books_stddev_samp_fields struct {
	Activities_count float64 `json:"activities_count"`
	Audio_seconds float64 `json:"audio_seconds"`
	Book_category_id float64 `json:"book_category_id"`
	Book_status_id float64 `json:"book_status_id"`
	Canonical_id float64 `json:"canonical_id"`
	Created_by_user_id float64 `json:"created_by_user_id"`
	Default_audio_edition_id float64 `json:"default_audio_edition_id"`
	Default_cover_edition_id float64 `json:"default_cover_edition_id"`
	Default_ebook_edition_id float64 `json:"default_ebook_edition_id"`
	Default_physical_edition_id float64 `json:"default_physical_edition_id"`
	Editions_count float64 `json:"editions_count"`
	Featured_book_series_id float64 `json:"featured_book_series_id"`
	Header_image_id float64 `json:"header_image_id"`
	ID float64 `json:"id"`
	Image_id float64 `json:"image_id"`
	Import_platform_id float64 `json:"import_platform_id"`
	Journals_count float64 `json:"journals_count"`
	Lists_count float64 `json:"lists_count"`
	Literary_type_id float64 `json:"literary_type_id"`
	Pages float64 `json:"pages"`
	Prompts_count float64 `json:"prompts_count"`
	Rating float64 `json:"rating"`
	Ratings_count float64 `json:"ratings_count"`
	Release_year float64 `json:"release_year"`
	Reviews_count float64 `json:"reviews_count"`
	Users_count float64 `json:"users_count"`
	Users_read_count float64 `json:"users_read_count"`
}

// books_stddev_samp_order_by represents the books_stddev_samp_order_by GraphQL type
type books_stddev_samp_order_by struct {
}

// books_stream_cursor_input represents the books_stream_cursor_input GraphQL type
type books_stream_cursor_input struct {
}

// books_stream_cursor_value_input represents the books_stream_cursor_value_input GraphQL type
type books_stream_cursor_value_input struct {
}

// books_sum_fields represents the books_sum_fields GraphQL type
type books_sum_fields struct {
	Activities_count int `json:"activities_count"`
	Audio_seconds int `json:"audio_seconds"`
	Book_category_id int `json:"book_category_id"`
	Book_status_id *smallint `json:"book_status_id"`
	Canonical_id int `json:"canonical_id"`
	Created_by_user_id int `json:"created_by_user_id"`
	Default_audio_edition_id int `json:"default_audio_edition_id"`
	Default_cover_edition_id int `json:"default_cover_edition_id"`
	Default_ebook_edition_id int `json:"default_ebook_edition_id"`
	Default_physical_edition_id int `json:"default_physical_edition_id"`
	Editions_count int `json:"editions_count"`
	Featured_book_series_id int `json:"featured_book_series_id"`
	Header_image_id int `json:"header_image_id"`
	ID int `json:"id"`
	Image_id int `json:"image_id"`
	Import_platform_id int `json:"import_platform_id"`
	Journals_count int `json:"journals_count"`
	Lists_count int `json:"lists_count"`
	Literary_type_id int `json:"literary_type_id"`
	Pages int `json:"pages"`
	Prompts_count int `json:"prompts_count"`
	Rating *numeric `json:"rating"`
	Ratings_count int `json:"ratings_count"`
	Release_year int `json:"release_year"`
	Reviews_count int `json:"reviews_count"`
	Users_count int `json:"users_count"`
	Users_read_count int `json:"users_read_count"`
}

// books_sum_order_by represents the books_sum_order_by GraphQL type
type books_sum_order_by struct {
}

// books_var_pop_fields represents the books_var_pop_fields GraphQL type
type books_var_pop_fields struct {
	Activities_count float64 `json:"activities_count"`
	Audio_seconds float64 `json:"audio_seconds"`
	Book_category_id float64 `json:"book_category_id"`
	Book_status_id float64 `json:"book_status_id"`
	Canonical_id float64 `json:"canonical_id"`
	Created_by_user_id float64 `json:"created_by_user_id"`
	Default_audio_edition_id float64 `json:"default_audio_edition_id"`
	Default_cover_edition_id float64 `json:"default_cover_edition_id"`
	Default_ebook_edition_id float64 `json:"default_ebook_edition_id"`
	Default_physical_edition_id float64 `json:"default_physical_edition_id"`
	Editions_count float64 `json:"editions_count"`
	Featured_book_series_id float64 `json:"featured_book_series_id"`
	Header_image_id float64 `json:"header_image_id"`
	ID float64 `json:"id"`
	Image_id float64 `json:"image_id"`
	Import_platform_id float64 `json:"import_platform_id"`
	Journals_count float64 `json:"journals_count"`
	Lists_count float64 `json:"lists_count"`
	Literary_type_id float64 `json:"literary_type_id"`
	Pages float64 `json:"pages"`
	Prompts_count float64 `json:"prompts_count"`
	Rating float64 `json:"rating"`
	Ratings_count float64 `json:"ratings_count"`
	Release_year float64 `json:"release_year"`
	Reviews_count float64 `json:"reviews_count"`
	Users_count float64 `json:"users_count"`
	Users_read_count float64 `json:"users_read_count"`
}

// books_var_pop_order_by represents the books_var_pop_order_by GraphQL type
type books_var_pop_order_by struct {
}

// books_var_samp_fields represents the books_var_samp_fields GraphQL type
type books_var_samp_fields struct {
	Activities_count float64 `json:"activities_count"`
	Audio_seconds float64 `json:"audio_seconds"`
	Book_category_id float64 `json:"book_category_id"`
	Book_status_id float64 `json:"book_status_id"`
	Canonical_id float64 `json:"canonical_id"`
	Created_by_user_id float64 `json:"created_by_user_id"`
	Default_audio_edition_id float64 `json:"default_audio_edition_id"`
	Default_cover_edition_id float64 `json:"default_cover_edition_id"`
	Default_ebook_edition_id float64 `json:"default_ebook_edition_id"`
	Default_physical_edition_id float64 `json:"default_physical_edition_id"`
	Editions_count float64 `json:"editions_count"`
	Featured_book_series_id float64 `json:"featured_book_series_id"`
	Header_image_id float64 `json:"header_image_id"`
	ID float64 `json:"id"`
	Image_id float64 `json:"image_id"`
	Import_platform_id float64 `json:"import_platform_id"`
	Journals_count float64 `json:"journals_count"`
	Lists_count float64 `json:"lists_count"`
	Literary_type_id float64 `json:"literary_type_id"`
	Pages float64 `json:"pages"`
	Prompts_count float64 `json:"prompts_count"`
	Rating float64 `json:"rating"`
	Ratings_count float64 `json:"ratings_count"`
	Release_year float64 `json:"release_year"`
	Reviews_count float64 `json:"reviews_count"`
	Users_count float64 `json:"users_count"`
	Users_read_count float64 `json:"users_read_count"`
}

// books_var_samp_order_by represents the books_var_samp_order_by GraphQL type
type books_var_samp_order_by struct {
}

// books_variance_fields represents the books_variance_fields GraphQL type
type books_variance_fields struct {
	Activities_count float64 `json:"activities_count"`
	Audio_seconds float64 `json:"audio_seconds"`
	Book_category_id float64 `json:"book_category_id"`
	Book_status_id float64 `json:"book_status_id"`
	Canonical_id float64 `json:"canonical_id"`
	Created_by_user_id float64 `json:"created_by_user_id"`
	Default_audio_edition_id float64 `json:"default_audio_edition_id"`
	Default_cover_edition_id float64 `json:"default_cover_edition_id"`
	Default_ebook_edition_id float64 `json:"default_ebook_edition_id"`
	Default_physical_edition_id float64 `json:"default_physical_edition_id"`
	Editions_count float64 `json:"editions_count"`
	Featured_book_series_id float64 `json:"featured_book_series_id"`
	Header_image_id float64 `json:"header_image_id"`
	ID float64 `json:"id"`
	Image_id float64 `json:"image_id"`
	Import_platform_id float64 `json:"import_platform_id"`
	Journals_count float64 `json:"journals_count"`
	Lists_count float64 `json:"lists_count"`
	Literary_type_id float64 `json:"literary_type_id"`
	Pages float64 `json:"pages"`
	Prompts_count float64 `json:"prompts_count"`
	Rating float64 `json:"rating"`
	Ratings_count float64 `json:"ratings_count"`
	Release_year float64 `json:"release_year"`
	Reviews_count float64 `json:"reviews_count"`
	Users_count float64 `json:"users_count"`
	Users_read_count float64 `json:"users_read_count"`
}

// books_variance_order_by represents the books_variance_order_by GraphQL type
type books_variance_order_by struct {
}

// characters represents the characters GraphQL type
type characters struct {
	Biography string `json:"biography"`
	Book_characters []*book_characters `json:"book_characters"`
	Books_count int `json:"books_count"`
	Cached_tags *json.RawMessage `json:"cached_tags"`
	Canonical *characters `json:"canonical"`
	Canonical_books_count int `json:"canonical_books_count"`
	Canonical_id int `json:"canonical_id"`
	Contributions []*contributions `json:"contributions"`
	Contributions_aggregate *contributions_aggregate `json:"contributions_aggregate"`
	Created_at *timestamp `json:"created_at"`
	Gender_id *bigint `json:"gender_id"`
	Has_disability bool `json:"has_disability"`
	ID *bigint `json:"id"`
	Image_id int `json:"image_id"`
	Is_lgbtq bool `json:"is_lgbtq"`
	Is_poc bool `json:"is_poc"`
	Locked bool `json:"locked"`
	Name string `json:"name"`
	Object_type string `json:"object_type"`
	Openlibrary_url string `json:"openlibrary_url"`
	Slug string `json:"slug"`
	State string `json:"state"`
	Updated_at *timestamp `json:"updated_at"`
	User_id int `json:"user_id"`
}

// characters_bool_exp represents the characters_bool_exp GraphQL type
type characters_bool_exp struct {
}

// characters_order_by represents the characters_order_by GraphQL type
type characters_order_by struct {
}

// characters_select_column represents the characters_select_column GraphQL type
type characters_select_column struct {
}

// characters_stream_cursor_input represents the characters_stream_cursor_input GraphQL type
type characters_stream_cursor_input struct {
}

// characters_stream_cursor_value_input represents the characters_stream_cursor_value_input GraphQL type
type characters_stream_cursor_value_input struct {
}

// citext represents the citext GraphQL type
type citext struct {
}

// citext_comparison_exp represents the citext_comparison_exp GraphQL type
type citext_comparison_exp struct {
}

// collection_import_results represents the collection_import_results GraphQL type
type collection_import_results struct {
	Author string `json:"author"`
	Book *books `json:"book"`
	Book_found_method string `json:"book_found_method"`
	Book_id int `json:"book_id"`
	Collection_import *collection_imports `json:"collection_import"`
	Collection_import_id int `json:"collection_import_id"`
	Contents *json.RawMessage `json:"contents"`
	External_id string `json:"external_id"`
	ID int `json:"id"`
	Report int `json:"report"`
	State string `json:"state"`
	Title string `json:"title"`
}

// collection_import_results_aggregate_order_by represents the collection_import_results_aggregate_order_by GraphQL type
type collection_import_results_aggregate_order_by struct {
}

// collection_import_results_avg_order_by represents the collection_import_results_avg_order_by GraphQL type
type collection_import_results_avg_order_by struct {
}

// collection_import_results_bool_exp represents the collection_import_results_bool_exp GraphQL type
type collection_import_results_bool_exp struct {
}

// collection_import_results_inc_input represents the collection_import_results_inc_input GraphQL type
type collection_import_results_inc_input struct {
}

// collection_import_results_max_order_by represents the collection_import_results_max_order_by GraphQL type
type collection_import_results_max_order_by struct {
}

// collection_import_results_min_order_by represents the collection_import_results_min_order_by GraphQL type
type collection_import_results_min_order_by struct {
}

// collection_import_results_mutation_response represents the collection_import_results_mutation_response GraphQL type
type collection_import_results_mutation_response struct {
	Affected_rows int `json:"affected_rows"`
	Returning []*collection_import_results `json:"returning"`
}

// collection_import_results_order_by represents the collection_import_results_order_by GraphQL type
type collection_import_results_order_by struct {
}

// collection_import_results_pk_columns_input represents the collection_import_results_pk_columns_input GraphQL type
type collection_import_results_pk_columns_input struct {
}

// collection_import_results_select_column represents the collection_import_results_select_column GraphQL type
type collection_import_results_select_column struct {
}

// collection_import_results_set_input represents the collection_import_results_set_input GraphQL type
type collection_import_results_set_input struct {
}

// collection_import_results_stddev_order_by represents the collection_import_results_stddev_order_by GraphQL type
type collection_import_results_stddev_order_by struct {
}

// collection_import_results_stddev_pop_order_by represents the collection_import_results_stddev_pop_order_by GraphQL type
type collection_import_results_stddev_pop_order_by struct {
}

// collection_import_results_stddev_samp_order_by represents the collection_import_results_stddev_samp_order_by GraphQL type
type collection_import_results_stddev_samp_order_by struct {
}

// collection_import_results_stream_cursor_input represents the collection_import_results_stream_cursor_input GraphQL type
type collection_import_results_stream_cursor_input struct {
}

// collection_import_results_stream_cursor_value_input represents the collection_import_results_stream_cursor_value_input GraphQL type
type collection_import_results_stream_cursor_value_input struct {
}

// collection_import_results_sum_order_by represents the collection_import_results_sum_order_by GraphQL type
type collection_import_results_sum_order_by struct {
}

// collection_import_results_updates represents the collection_import_results_updates GraphQL type
type collection_import_results_updates struct {
}

// collection_import_results_var_pop_order_by represents the collection_import_results_var_pop_order_by GraphQL type
type collection_import_results_var_pop_order_by struct {
}

// collection_import_results_var_samp_order_by represents the collection_import_results_var_samp_order_by GraphQL type
type collection_import_results_var_samp_order_by struct {
}

// collection_import_results_variance_order_by represents the collection_import_results_variance_order_by GraphQL type
type collection_import_results_variance_order_by struct {
}

// collection_imports represents the collection_imports GraphQL type
type collection_imports struct {
	Collection_import_results []*collection_import_results `json:"collection_import_results"`
	Completed_at *timestamptz `json:"completed_at"`
	Contents_key string `json:"contents_key"`
	Created_at *timestamptz `json:"created_at"`
	Current_book string `json:"current_book"`
	Error_message string `json:"error_message"`
	Failure_count int `json:"failure_count"`
	ID int `json:"id"`
	Override_date_read bool `json:"override_date_read"`
	Override_ratings bool `json:"override_ratings"`
	Override_shelves bool `json:"override_shelves"`
	Platform_id int `json:"platform_id"`
	Processed_count int `json:"processed_count"`
	Reimport_count int `json:"reimport_count"`
	Started_at *timestamptz `json:"started_at"`
	State string `json:"state"`
	Success_count int `json:"success_count"`
	Tag_resolution int `json:"tag_resolution"`
	Total_count int `json:"total_count"`
	Updated_at *timestamptz `json:"updated_at"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// collection_imports_aggregate_order_by represents the collection_imports_aggregate_order_by GraphQL type
type collection_imports_aggregate_order_by struct {
}

// collection_imports_avg_order_by represents the collection_imports_avg_order_by GraphQL type
type collection_imports_avg_order_by struct {
}

// collection_imports_bool_exp represents the collection_imports_bool_exp GraphQL type
type collection_imports_bool_exp struct {
}

// collection_imports_max_order_by represents the collection_imports_max_order_by GraphQL type
type collection_imports_max_order_by struct {
}

// collection_imports_min_order_by represents the collection_imports_min_order_by GraphQL type
type collection_imports_min_order_by struct {
}

// collection_imports_order_by represents the collection_imports_order_by GraphQL type
type collection_imports_order_by struct {
}

// collection_imports_select_column represents the collection_imports_select_column GraphQL type
type collection_imports_select_column struct {
}

// collection_imports_stddev_order_by represents the collection_imports_stddev_order_by GraphQL type
type collection_imports_stddev_order_by struct {
}

// collection_imports_stddev_pop_order_by represents the collection_imports_stddev_pop_order_by GraphQL type
type collection_imports_stddev_pop_order_by struct {
}

// collection_imports_stddev_samp_order_by represents the collection_imports_stddev_samp_order_by GraphQL type
type collection_imports_stddev_samp_order_by struct {
}

// collection_imports_stream_cursor_input represents the collection_imports_stream_cursor_input GraphQL type
type collection_imports_stream_cursor_input struct {
}

// collection_imports_stream_cursor_value_input represents the collection_imports_stream_cursor_value_input GraphQL type
type collection_imports_stream_cursor_value_input struct {
}

// collection_imports_sum_order_by represents the collection_imports_sum_order_by GraphQL type
type collection_imports_sum_order_by struct {
}

// collection_imports_var_pop_order_by represents the collection_imports_var_pop_order_by GraphQL type
type collection_imports_var_pop_order_by struct {
}

// collection_imports_var_samp_order_by represents the collection_imports_var_samp_order_by GraphQL type
type collection_imports_var_samp_order_by struct {
}

// collection_imports_variance_order_by represents the collection_imports_variance_order_by GraphQL type
type collection_imports_variance_order_by struct {
}

// contributions represents the contributions GraphQL type
type contributions struct {
	Author *authors `json:"author"`
	Author_id int `json:"author_id"`
	Book *books `json:"book"`
	Contributable_id int `json:"contributable_id"`
	Contributable_type string `json:"contributable_type"`
	Contribution string `json:"contribution"`
	Created_at *timestamp `json:"created_at"`
	ID *bigint `json:"id"`
	Updated_at *timestamp `json:"updated_at"`
}

// contributions_aggregate represents the contributions_aggregate GraphQL type
type contributions_aggregate struct {
	Aggregate *contributions_aggregate_fields `json:"aggregate"`
	Nodes []*contributions `json:"nodes"`
}

// contributions_aggregate_bool_exp represents the contributions_aggregate_bool_exp GraphQL type
type contributions_aggregate_bool_exp struct {
}

// contributions_aggregate_bool_exp_count represents the contributions_aggregate_bool_exp_count GraphQL type
type contributions_aggregate_bool_exp_count struct {
}

// contributions_aggregate_fields represents the contributions_aggregate_fields GraphQL type
type contributions_aggregate_fields struct {
	Avg *contributions_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *contributions_max_fields `json:"max"`
	Min *contributions_min_fields `json:"min"`
	Stddev *contributions_stddev_fields `json:"stddev"`
	Stddev_pop *contributions_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *contributions_stddev_samp_fields `json:"stddev_samp"`
	Sum *contributions_sum_fields `json:"sum"`
	Var_pop *contributions_var_pop_fields `json:"var_pop"`
	Var_samp *contributions_var_samp_fields `json:"var_samp"`
	Variance *contributions_variance_fields `json:"variance"`
}

// contributions_aggregate_order_by represents the contributions_aggregate_order_by GraphQL type
type contributions_aggregate_order_by struct {
}

// contributions_avg_fields represents the contributions_avg_fields GraphQL type
type contributions_avg_fields struct {
	Author_id float64 `json:"author_id"`
	Contributable_id float64 `json:"contributable_id"`
	ID float64 `json:"id"`
}

// contributions_avg_order_by represents the contributions_avg_order_by GraphQL type
type contributions_avg_order_by struct {
}

// contributions_bool_exp represents the contributions_bool_exp GraphQL type
type contributions_bool_exp struct {
}

// contributions_max_fields represents the contributions_max_fields GraphQL type
type contributions_max_fields struct {
	Author_id int `json:"author_id"`
	Contributable_id int `json:"contributable_id"`
	Contributable_type string `json:"contributable_type"`
	Contribution string `json:"contribution"`
	Created_at *timestamp `json:"created_at"`
	ID *bigint `json:"id"`
	Updated_at *timestamp `json:"updated_at"`
}

// contributions_max_order_by represents the contributions_max_order_by GraphQL type
type contributions_max_order_by struct {
}

// contributions_min_fields represents the contributions_min_fields GraphQL type
type contributions_min_fields struct {
	Author_id int `json:"author_id"`
	Contributable_id int `json:"contributable_id"`
	Contributable_type string `json:"contributable_type"`
	Contribution string `json:"contribution"`
	Created_at *timestamp `json:"created_at"`
	ID *bigint `json:"id"`
	Updated_at *timestamp `json:"updated_at"`
}

// contributions_min_order_by represents the contributions_min_order_by GraphQL type
type contributions_min_order_by struct {
}

// contributions_order_by represents the contributions_order_by GraphQL type
type contributions_order_by struct {
}

// contributions_select_column represents the contributions_select_column GraphQL type
type contributions_select_column struct {
}

// contributions_stddev_fields represents the contributions_stddev_fields GraphQL type
type contributions_stddev_fields struct {
	Author_id float64 `json:"author_id"`
	Contributable_id float64 `json:"contributable_id"`
	ID float64 `json:"id"`
}

// contributions_stddev_order_by represents the contributions_stddev_order_by GraphQL type
type contributions_stddev_order_by struct {
}

// contributions_stddev_pop_fields represents the contributions_stddev_pop_fields GraphQL type
type contributions_stddev_pop_fields struct {
	Author_id float64 `json:"author_id"`
	Contributable_id float64 `json:"contributable_id"`
	ID float64 `json:"id"`
}

// contributions_stddev_pop_order_by represents the contributions_stddev_pop_order_by GraphQL type
type contributions_stddev_pop_order_by struct {
}

// contributions_stddev_samp_fields represents the contributions_stddev_samp_fields GraphQL type
type contributions_stddev_samp_fields struct {
	Author_id float64 `json:"author_id"`
	Contributable_id float64 `json:"contributable_id"`
	ID float64 `json:"id"`
}

// contributions_stddev_samp_order_by represents the contributions_stddev_samp_order_by GraphQL type
type contributions_stddev_samp_order_by struct {
}

// contributions_stream_cursor_input represents the contributions_stream_cursor_input GraphQL type
type contributions_stream_cursor_input struct {
}

// contributions_stream_cursor_value_input represents the contributions_stream_cursor_value_input GraphQL type
type contributions_stream_cursor_value_input struct {
}

// contributions_sum_fields represents the contributions_sum_fields GraphQL type
type contributions_sum_fields struct {
	Author_id int `json:"author_id"`
	Contributable_id int `json:"contributable_id"`
	ID *bigint `json:"id"`
}

// contributions_sum_order_by represents the contributions_sum_order_by GraphQL type
type contributions_sum_order_by struct {
}

// contributions_var_pop_fields represents the contributions_var_pop_fields GraphQL type
type contributions_var_pop_fields struct {
	Author_id float64 `json:"author_id"`
	Contributable_id float64 `json:"contributable_id"`
	ID float64 `json:"id"`
}

// contributions_var_pop_order_by represents the contributions_var_pop_order_by GraphQL type
type contributions_var_pop_order_by struct {
}

// contributions_var_samp_fields represents the contributions_var_samp_fields GraphQL type
type contributions_var_samp_fields struct {
	Author_id float64 `json:"author_id"`
	Contributable_id float64 `json:"contributable_id"`
	ID float64 `json:"id"`
}

// contributions_var_samp_order_by represents the contributions_var_samp_order_by GraphQL type
type contributions_var_samp_order_by struct {
}

// contributions_variance_fields represents the contributions_variance_fields GraphQL type
type contributions_variance_fields struct {
	Author_id float64 `json:"author_id"`
	Contributable_id float64 `json:"contributable_id"`
	ID float64 `json:"id"`
}

// contributions_variance_order_by represents the contributions_variance_order_by GraphQL type
type contributions_variance_order_by struct {
}

// countries represents the countries GraphQL type
type countries struct {
	Code2 string `json:"code2"`
	Code3 string `json:"code3"`
	Created_at *timestamp `json:"created_at"`
	Editions []*editions `json:"editions"`
	ID *bigint `json:"id"`
	Intermediate_region string `json:"intermediate_region"`
	Intermediate_region_code string `json:"intermediate_region_code"`
	Iso_3166 string `json:"iso_3166"`
	Name string `json:"name"`
	Phone_code string `json:"phone_code"`
	Region string `json:"region"`
	Region_code string `json:"region_code"`
	Sub_region string `json:"sub_region"`
	Sub_region_code string `json:"sub_region_code"`
	Updated_at *timestamp `json:"updated_at"`
}

// countries_bool_exp represents the countries_bool_exp GraphQL type
type countries_bool_exp struct {
}

// countries_order_by represents the countries_order_by GraphQL type
type countries_order_by struct {
}

// countries_select_column represents the countries_select_column GraphQL type
type countries_select_column struct {
}

// countries_stream_cursor_input represents the countries_stream_cursor_input GraphQL type
type countries_stream_cursor_input struct {
}

// countries_stream_cursor_value_input represents the countries_stream_cursor_value_input GraphQL type
type countries_stream_cursor_value_input struct {
}

// cursor_ordering represents the cursor_ordering GraphQL type
type cursor_ordering struct {
}

// date represents the date GraphQL type
type date struct {
}

// date_comparison_exp represents the date_comparison_exp GraphQL type
type date_comparison_exp struct {
}

// editions represents the editions GraphQL type
type editions struct {
	Alternative_titles *json.RawMessage `json:"alternative_titles"`
	Asin string `json:"asin"`
	Audio_seconds int `json:"audio_seconds"`
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Book_mappings []*book_mappings `json:"book_mappings"`
	Cached_contributors *json.RawMessage `json:"cached_contributors"`
	Cached_image *json.RawMessage `json:"cached_image"`
	Cached_tags *json.RawMessage `json:"cached_tags"`
	Compilation bool `json:"compilation"`
	Contributions []*contributions `json:"contributions"`
	Contributions_aggregate *contributions_aggregate `json:"contributions_aggregate"`
	Country *countries `json:"country"`
	Country_id int `json:"country_id"`
	Created_at *timestamp `json:"created_at"`
	Created_by_user_id int `json:"created_by_user_id"`
	Dto *json.RawMessage `json:"dto"`
	Dto_combined *json.RawMessage `json:"dto_combined"`
	Dto_external *json.RawMessage `json:"dto_external"`
	Edition_format string `json:"edition_format"`
	Edition_information string `json:"edition_information"`
	ID int `json:"id"`
	Image *images `json:"image"`
	Image_id int `json:"image_id"`
	Images []*images `json:"images"`
	Isbn_10 string `json:"isbn_10"`
	Isbn_13 string `json:"isbn_13"`
	Language *languages `json:"language"`
	Language_id int `json:"language_id"`
	List_books []*list_books `json:"list_books"`
	List_books_aggregate *list_books_aggregate `json:"list_books_aggregate"`
	Lists_count int `json:"lists_count"`
	Locked bool `json:"locked"`
	Normalized_at *timestamp `json:"normalized_at"`
	Object_type string `json:"object_type"`
	Original_book_id int `json:"original_book_id"`
	Pages int `json:"pages"`
	Physical_format string `json:"physical_format"`
	Physical_information string `json:"physical_information"`
	Publisher *publishers `json:"publisher"`
	Publisher_id int `json:"publisher_id"`
	Rating *numeric `json:"rating"`
	Reading_format *reading_formats `json:"reading_format"`
	Reading_format_id int `json:"reading_format_id"`
	Release_date *date `json:"release_date"`
	Release_year int `json:"release_year"`
	Score int `json:"score"`
	Source string `json:"source"`
	State string `json:"state"`
	Subtitle string `json:"subtitle"`
	Title string `json:"title"`
	Updated_at *timestamp `json:"updated_at"`
	User_added bool `json:"user_added"`
	Users_count int `json:"users_count"`
	Users_read_count int `json:"users_read_count"`
}

// editions_aggregate_order_by represents the editions_aggregate_order_by GraphQL type
type editions_aggregate_order_by struct {
}

// editions_avg_order_by represents the editions_avg_order_by GraphQL type
type editions_avg_order_by struct {
}

// editions_bool_exp represents the editions_bool_exp GraphQL type
type editions_bool_exp struct {
}

// editions_max_order_by represents the editions_max_order_by GraphQL type
type editions_max_order_by struct {
}

// editions_min_order_by represents the editions_min_order_by GraphQL type
type editions_min_order_by struct {
}

// editions_order_by represents the editions_order_by GraphQL type
type editions_order_by struct {
}

// editions_select_column represents the editions_select_column GraphQL type
type editions_select_column struct {
}

// editions_stddev_order_by represents the editions_stddev_order_by GraphQL type
type editions_stddev_order_by struct {
}

// editions_stddev_pop_order_by represents the editions_stddev_pop_order_by GraphQL type
type editions_stddev_pop_order_by struct {
}

// editions_stddev_samp_order_by represents the editions_stddev_samp_order_by GraphQL type
type editions_stddev_samp_order_by struct {
}

// editions_stream_cursor_input represents the editions_stream_cursor_input GraphQL type
type editions_stream_cursor_input struct {
}

// editions_stream_cursor_value_input represents the editions_stream_cursor_value_input GraphQL type
type editions_stream_cursor_value_input struct {
}

// editions_sum_order_by represents the editions_sum_order_by GraphQL type
type editions_sum_order_by struct {
}

// editions_var_pop_order_by represents the editions_var_pop_order_by GraphQL type
type editions_var_pop_order_by struct {
}

// editions_var_samp_order_by represents the editions_var_samp_order_by GraphQL type
type editions_var_samp_order_by struct {
}

// editions_variance_order_by represents the editions_variance_order_by GraphQL type
type editions_variance_order_by struct {
}

// flag_statuses represents the flag_statuses GraphQL type
type flag_statuses struct {
	ID int `json:"id"`
	Status string `json:"status"`
	User_flags []*user_flags `json:"user_flags"`
}

// flag_statuses_bool_exp represents the flag_statuses_bool_exp GraphQL type
type flag_statuses_bool_exp struct {
}

// flag_statuses_order_by represents the flag_statuses_order_by GraphQL type
type flag_statuses_order_by struct {
}

// flag_statuses_select_column represents the flag_statuses_select_column GraphQL type
type flag_statuses_select_column struct {
}

// flag_statuses_stream_cursor_input represents the flag_statuses_stream_cursor_input GraphQL type
type flag_statuses_stream_cursor_input struct {
}

// flag_statuses_stream_cursor_value_input represents the flag_statuses_stream_cursor_value_input GraphQL type
type flag_statuses_stream_cursor_value_input struct {
}

// float8 represents the float8 GraphQL type
type float8 struct {
}

// float8_comparison_exp represents the float8_comparison_exp GraphQL type
type float8_comparison_exp struct {
}

// followed_lists represents the followed_lists GraphQL type
type followed_lists struct {
	Created_at *timestamptz `json:"created_at"`
	ID int `json:"id"`
	List *lists `json:"list"`
	List_id int `json:"list_id"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// followed_lists_aggregate_order_by represents the followed_lists_aggregate_order_by GraphQL type
type followed_lists_aggregate_order_by struct {
}

// followed_lists_avg_order_by represents the followed_lists_avg_order_by GraphQL type
type followed_lists_avg_order_by struct {
}

// followed_lists_bool_exp represents the followed_lists_bool_exp GraphQL type
type followed_lists_bool_exp struct {
}

// followed_lists_max_order_by represents the followed_lists_max_order_by GraphQL type
type followed_lists_max_order_by struct {
}

// followed_lists_min_order_by represents the followed_lists_min_order_by GraphQL type
type followed_lists_min_order_by struct {
}

// followed_lists_order_by represents the followed_lists_order_by GraphQL type
type followed_lists_order_by struct {
}

// followed_lists_select_column represents the followed_lists_select_column GraphQL type
type followed_lists_select_column struct {
}

// followed_lists_stddev_order_by represents the followed_lists_stddev_order_by GraphQL type
type followed_lists_stddev_order_by struct {
}

// followed_lists_stddev_pop_order_by represents the followed_lists_stddev_pop_order_by GraphQL type
type followed_lists_stddev_pop_order_by struct {
}

// followed_lists_stddev_samp_order_by represents the followed_lists_stddev_samp_order_by GraphQL type
type followed_lists_stddev_samp_order_by struct {
}

// followed_lists_stream_cursor_input represents the followed_lists_stream_cursor_input GraphQL type
type followed_lists_stream_cursor_input struct {
}

// followed_lists_stream_cursor_value_input represents the followed_lists_stream_cursor_value_input GraphQL type
type followed_lists_stream_cursor_value_input struct {
}

// followed_lists_sum_order_by represents the followed_lists_sum_order_by GraphQL type
type followed_lists_sum_order_by struct {
}

// followed_lists_var_pop_order_by represents the followed_lists_var_pop_order_by GraphQL type
type followed_lists_var_pop_order_by struct {
}

// followed_lists_var_samp_order_by represents the followed_lists_var_samp_order_by GraphQL type
type followed_lists_var_samp_order_by struct {
}

// followed_lists_variance_order_by represents the followed_lists_variance_order_by GraphQL type
type followed_lists_variance_order_by struct {
}

// followed_prompts represents the followed_prompts GraphQL type
type followed_prompts struct {
	Created_at *timestamptz `json:"created_at"`
	ID int `json:"id"`
	Order int `json:"order"`
	Prompt *prompts `json:"prompt"`
	Prompt_id int `json:"prompt_id"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// followed_prompts_aggregate_order_by represents the followed_prompts_aggregate_order_by GraphQL type
type followed_prompts_aggregate_order_by struct {
}

// followed_prompts_avg_order_by represents the followed_prompts_avg_order_by GraphQL type
type followed_prompts_avg_order_by struct {
}

// followed_prompts_bool_exp represents the followed_prompts_bool_exp GraphQL type
type followed_prompts_bool_exp struct {
}

// followed_prompts_constraint represents the followed_prompts_constraint GraphQL type
type followed_prompts_constraint struct {
}

// followed_prompts_inc_input represents the followed_prompts_inc_input GraphQL type
type followed_prompts_inc_input struct {
}

// followed_prompts_insert_input represents the followed_prompts_insert_input GraphQL type
type followed_prompts_insert_input struct {
}

// followed_prompts_max_order_by represents the followed_prompts_max_order_by GraphQL type
type followed_prompts_max_order_by struct {
}

// followed_prompts_min_order_by represents the followed_prompts_min_order_by GraphQL type
type followed_prompts_min_order_by struct {
}

// followed_prompts_mutation_response represents the followed_prompts_mutation_response GraphQL type
type followed_prompts_mutation_response struct {
	Affected_rows int `json:"affected_rows"`
	Returning []*followed_prompts `json:"returning"`
}

// followed_prompts_on_conflict represents the followed_prompts_on_conflict GraphQL type
type followed_prompts_on_conflict struct {
}

// followed_prompts_order_by represents the followed_prompts_order_by GraphQL type
type followed_prompts_order_by struct {
}

// followed_prompts_pk_columns_input represents the followed_prompts_pk_columns_input GraphQL type
type followed_prompts_pk_columns_input struct {
}

// followed_prompts_select_column represents the followed_prompts_select_column GraphQL type
type followed_prompts_select_column struct {
}

// followed_prompts_set_input represents the followed_prompts_set_input GraphQL type
type followed_prompts_set_input struct {
}

// followed_prompts_stddev_order_by represents the followed_prompts_stddev_order_by GraphQL type
type followed_prompts_stddev_order_by struct {
}

// followed_prompts_stddev_pop_order_by represents the followed_prompts_stddev_pop_order_by GraphQL type
type followed_prompts_stddev_pop_order_by struct {
}

// followed_prompts_stddev_samp_order_by represents the followed_prompts_stddev_samp_order_by GraphQL type
type followed_prompts_stddev_samp_order_by struct {
}

// followed_prompts_stream_cursor_input represents the followed_prompts_stream_cursor_input GraphQL type
type followed_prompts_stream_cursor_input struct {
}

// followed_prompts_stream_cursor_value_input represents the followed_prompts_stream_cursor_value_input GraphQL type
type followed_prompts_stream_cursor_value_input struct {
}

// followed_prompts_sum_order_by represents the followed_prompts_sum_order_by GraphQL type
type followed_prompts_sum_order_by struct {
}

// followed_prompts_update_column represents the followed_prompts_update_column GraphQL type
type followed_prompts_update_column struct {
}

// followed_prompts_updates represents the followed_prompts_updates GraphQL type
type followed_prompts_updates struct {
}

// followed_prompts_var_pop_order_by represents the followed_prompts_var_pop_order_by GraphQL type
type followed_prompts_var_pop_order_by struct {
}

// followed_prompts_var_samp_order_by represents the followed_prompts_var_samp_order_by GraphQL type
type followed_prompts_var_samp_order_by struct {
}

// followed_prompts_variance_order_by represents the followed_prompts_variance_order_by GraphQL type
type followed_prompts_variance_order_by struct {
}

// followed_user_books represents the followed_user_books GraphQL type
type followed_user_books struct {
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Follower_user *users `json:"follower_user"`
	Follower_user_id int `json:"follower_user_id"`
	User *users `json:"user"`
	User_book *user_books `json:"user_book"`
	User_book_id int `json:"user_book_id"`
	User_id int `json:"user_id"`
}

// followed_user_books_aggregate represents the followed_user_books_aggregate GraphQL type
type followed_user_books_aggregate struct {
	Aggregate *followed_user_books_aggregate_fields `json:"aggregate"`
	Nodes []*followed_user_books `json:"nodes"`
}

// followed_user_books_aggregate_fields represents the followed_user_books_aggregate_fields GraphQL type
type followed_user_books_aggregate_fields struct {
	Avg *followed_user_books_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *followed_user_books_max_fields `json:"max"`
	Min *followed_user_books_min_fields `json:"min"`
	Stddev *followed_user_books_stddev_fields `json:"stddev"`
	Stddev_pop *followed_user_books_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *followed_user_books_stddev_samp_fields `json:"stddev_samp"`
	Sum *followed_user_books_sum_fields `json:"sum"`
	Var_pop *followed_user_books_var_pop_fields `json:"var_pop"`
	Var_samp *followed_user_books_var_samp_fields `json:"var_samp"`
	Variance *followed_user_books_variance_fields `json:"variance"`
}

// followed_user_books_avg_fields represents the followed_user_books_avg_fields GraphQL type
type followed_user_books_avg_fields struct {
	Book_id float64 `json:"book_id"`
	Follower_user_id float64 `json:"follower_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// followed_user_books_bool_exp represents the followed_user_books_bool_exp GraphQL type
type followed_user_books_bool_exp struct {
}

// followed_user_books_max_fields represents the followed_user_books_max_fields GraphQL type
type followed_user_books_max_fields struct {
	Book_id int `json:"book_id"`
	Follower_user_id int `json:"follower_user_id"`
	User_book_id int `json:"user_book_id"`
	User_id int `json:"user_id"`
}

// followed_user_books_min_fields represents the followed_user_books_min_fields GraphQL type
type followed_user_books_min_fields struct {
	Book_id int `json:"book_id"`
	Follower_user_id int `json:"follower_user_id"`
	User_book_id int `json:"user_book_id"`
	User_id int `json:"user_id"`
}

// followed_user_books_order_by represents the followed_user_books_order_by GraphQL type
type followed_user_books_order_by struct {
}

// followed_user_books_select_column represents the followed_user_books_select_column GraphQL type
type followed_user_books_select_column struct {
}

// followed_user_books_stddev_fields represents the followed_user_books_stddev_fields GraphQL type
type followed_user_books_stddev_fields struct {
	Book_id float64 `json:"book_id"`
	Follower_user_id float64 `json:"follower_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// followed_user_books_stddev_pop_fields represents the followed_user_books_stddev_pop_fields GraphQL type
type followed_user_books_stddev_pop_fields struct {
	Book_id float64 `json:"book_id"`
	Follower_user_id float64 `json:"follower_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// followed_user_books_stddev_samp_fields represents the followed_user_books_stddev_samp_fields GraphQL type
type followed_user_books_stddev_samp_fields struct {
	Book_id float64 `json:"book_id"`
	Follower_user_id float64 `json:"follower_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// followed_user_books_stream_cursor_input represents the followed_user_books_stream_cursor_input GraphQL type
type followed_user_books_stream_cursor_input struct {
}

// followed_user_books_stream_cursor_value_input represents the followed_user_books_stream_cursor_value_input GraphQL type
type followed_user_books_stream_cursor_value_input struct {
}

// followed_user_books_sum_fields represents the followed_user_books_sum_fields GraphQL type
type followed_user_books_sum_fields struct {
	Book_id int `json:"book_id"`
	Follower_user_id int `json:"follower_user_id"`
	User_book_id int `json:"user_book_id"`
	User_id int `json:"user_id"`
}

// followed_user_books_var_pop_fields represents the followed_user_books_var_pop_fields GraphQL type
type followed_user_books_var_pop_fields struct {
	Book_id float64 `json:"book_id"`
	Follower_user_id float64 `json:"follower_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// followed_user_books_var_samp_fields represents the followed_user_books_var_samp_fields GraphQL type
type followed_user_books_var_samp_fields struct {
	Book_id float64 `json:"book_id"`
	Follower_user_id float64 `json:"follower_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// followed_user_books_variance_fields represents the followed_user_books_variance_fields GraphQL type
type followed_user_books_variance_fields struct {
	Book_id float64 `json:"book_id"`
	Follower_user_id float64 `json:"follower_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// followed_users represents the followed_users GraphQL type
type followed_users struct {
	Created_at *timestamptz `json:"created_at"`
	Followed_user *users `json:"followed_user"`
	Followed_user_id int `json:"followed_user_id"`
	ID int `json:"id"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// followed_users_aggregate_order_by represents the followed_users_aggregate_order_by GraphQL type
type followed_users_aggregate_order_by struct {
}

// followed_users_avg_order_by represents the followed_users_avg_order_by GraphQL type
type followed_users_avg_order_by struct {
}

// followed_users_bool_exp represents the followed_users_bool_exp GraphQL type
type followed_users_bool_exp struct {
}

// followed_users_max_order_by represents the followed_users_max_order_by GraphQL type
type followed_users_max_order_by struct {
}

// followed_users_min_order_by represents the followed_users_min_order_by GraphQL type
type followed_users_min_order_by struct {
}

// followed_users_mutation_response represents the followed_users_mutation_response GraphQL type
type followed_users_mutation_response struct {
	Affected_rows int `json:"affected_rows"`
	Returning []*followed_users `json:"returning"`
}

// followed_users_order_by represents the followed_users_order_by GraphQL type
type followed_users_order_by struct {
}

// followed_users_select_column represents the followed_users_select_column GraphQL type
type followed_users_select_column struct {
}

// followed_users_stddev_order_by represents the followed_users_stddev_order_by GraphQL type
type followed_users_stddev_order_by struct {
}

// followed_users_stddev_pop_order_by represents the followed_users_stddev_pop_order_by GraphQL type
type followed_users_stddev_pop_order_by struct {
}

// followed_users_stddev_samp_order_by represents the followed_users_stddev_samp_order_by GraphQL type
type followed_users_stddev_samp_order_by struct {
}

// followed_users_stream_cursor_input represents the followed_users_stream_cursor_input GraphQL type
type followed_users_stream_cursor_input struct {
}

// followed_users_stream_cursor_value_input represents the followed_users_stream_cursor_value_input GraphQL type
type followed_users_stream_cursor_value_input struct {
}

// followed_users_sum_order_by represents the followed_users_sum_order_by GraphQL type
type followed_users_sum_order_by struct {
}

// followed_users_var_pop_order_by represents the followed_users_var_pop_order_by GraphQL type
type followed_users_var_pop_order_by struct {
}

// followed_users_var_samp_order_by represents the followed_users_var_samp_order_by GraphQL type
type followed_users_var_samp_order_by struct {
}

// followed_users_variance_order_by represents the followed_users_variance_order_by GraphQL type
type followed_users_variance_order_by struct {
}

// following_user_books represents the following_user_books GraphQL type
type following_user_books struct {
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Followed_user_id int `json:"followed_user_id"`
	Following_user *users `json:"following_user"`
	User *users `json:"user"`
	User_book *user_books `json:"user_book"`
	User_book_id int `json:"user_book_id"`
	User_id int `json:"user_id"`
}

// following_user_books_aggregate represents the following_user_books_aggregate GraphQL type
type following_user_books_aggregate struct {
	Aggregate *following_user_books_aggregate_fields `json:"aggregate"`
	Nodes []*following_user_books `json:"nodes"`
}

// following_user_books_aggregate_fields represents the following_user_books_aggregate_fields GraphQL type
type following_user_books_aggregate_fields struct {
	Avg *following_user_books_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *following_user_books_max_fields `json:"max"`
	Min *following_user_books_min_fields `json:"min"`
	Stddev *following_user_books_stddev_fields `json:"stddev"`
	Stddev_pop *following_user_books_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *following_user_books_stddev_samp_fields `json:"stddev_samp"`
	Sum *following_user_books_sum_fields `json:"sum"`
	Var_pop *following_user_books_var_pop_fields `json:"var_pop"`
	Var_samp *following_user_books_var_samp_fields `json:"var_samp"`
	Variance *following_user_books_variance_fields `json:"variance"`
}

// following_user_books_avg_fields represents the following_user_books_avg_fields GraphQL type
type following_user_books_avg_fields struct {
	Book_id float64 `json:"book_id"`
	Followed_user_id float64 `json:"followed_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// following_user_books_bool_exp represents the following_user_books_bool_exp GraphQL type
type following_user_books_bool_exp struct {
}

// following_user_books_max_fields represents the following_user_books_max_fields GraphQL type
type following_user_books_max_fields struct {
	Book_id int `json:"book_id"`
	Followed_user_id int `json:"followed_user_id"`
	User_book_id int `json:"user_book_id"`
	User_id int `json:"user_id"`
}

// following_user_books_min_fields represents the following_user_books_min_fields GraphQL type
type following_user_books_min_fields struct {
	Book_id int `json:"book_id"`
	Followed_user_id int `json:"followed_user_id"`
	User_book_id int `json:"user_book_id"`
	User_id int `json:"user_id"`
}

// following_user_books_order_by represents the following_user_books_order_by GraphQL type
type following_user_books_order_by struct {
}

// following_user_books_select_column represents the following_user_books_select_column GraphQL type
type following_user_books_select_column struct {
}

// following_user_books_stddev_fields represents the following_user_books_stddev_fields GraphQL type
type following_user_books_stddev_fields struct {
	Book_id float64 `json:"book_id"`
	Followed_user_id float64 `json:"followed_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// following_user_books_stddev_pop_fields represents the following_user_books_stddev_pop_fields GraphQL type
type following_user_books_stddev_pop_fields struct {
	Book_id float64 `json:"book_id"`
	Followed_user_id float64 `json:"followed_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// following_user_books_stddev_samp_fields represents the following_user_books_stddev_samp_fields GraphQL type
type following_user_books_stddev_samp_fields struct {
	Book_id float64 `json:"book_id"`
	Followed_user_id float64 `json:"followed_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// following_user_books_stream_cursor_input represents the following_user_books_stream_cursor_input GraphQL type
type following_user_books_stream_cursor_input struct {
}

// following_user_books_stream_cursor_value_input represents the following_user_books_stream_cursor_value_input GraphQL type
type following_user_books_stream_cursor_value_input struct {
}

// following_user_books_sum_fields represents the following_user_books_sum_fields GraphQL type
type following_user_books_sum_fields struct {
	Book_id int `json:"book_id"`
	Followed_user_id int `json:"followed_user_id"`
	User_book_id int `json:"user_book_id"`
	User_id int `json:"user_id"`
}

// following_user_books_var_pop_fields represents the following_user_books_var_pop_fields GraphQL type
type following_user_books_var_pop_fields struct {
	Book_id float64 `json:"book_id"`
	Followed_user_id float64 `json:"followed_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// following_user_books_var_samp_fields represents the following_user_books_var_samp_fields GraphQL type
type following_user_books_var_samp_fields struct {
	Book_id float64 `json:"book_id"`
	Followed_user_id float64 `json:"followed_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// following_user_books_variance_fields represents the following_user_books_variance_fields GraphQL type
type following_user_books_variance_fields struct {
	Book_id float64 `json:"book_id"`
	Followed_user_id float64 `json:"followed_user_id"`
	User_book_id float64 `json:"user_book_id"`
	User_id float64 `json:"user_id"`
}

// goals represents the goals GraphQL type
type goals struct {
	Archived bool `json:"archived"`
	Completed_at *timestamptz `json:"completed_at"`
	Conditions *json.RawMessage `json:"conditions"`
	Description string `json:"description"`
	End_date *date `json:"end_date"`
	Followers []*followed_users `json:"followers"`
	Goal int `json:"goal"`
	ID int `json:"id"`
	Metric string `json:"metric"`
	Privacy_setting_id int `json:"privacy_setting_id"`
	Progress *numeric `json:"progress"`
	Start_date *date `json:"start_date"`
	State string `json:"state"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// goals_aggregate_order_by represents the goals_aggregate_order_by GraphQL type
type goals_aggregate_order_by struct {
}

// goals_avg_order_by represents the goals_avg_order_by GraphQL type
type goals_avg_order_by struct {
}

// goals_bool_exp represents the goals_bool_exp GraphQL type
type goals_bool_exp struct {
}

// goals_max_order_by represents the goals_max_order_by GraphQL type
type goals_max_order_by struct {
}

// goals_min_order_by represents the goals_min_order_by GraphQL type
type goals_min_order_by struct {
}

// goals_mutation_response represents the goals_mutation_response GraphQL type
type goals_mutation_response struct {
	Affected_rows int `json:"affected_rows"`
	Returning []*goals `json:"returning"`
}

// goals_order_by represents the goals_order_by GraphQL type
type goals_order_by struct {
}

// goals_select_column represents the goals_select_column GraphQL type
type goals_select_column struct {
}

// goals_stddev_order_by represents the goals_stddev_order_by GraphQL type
type goals_stddev_order_by struct {
}

// goals_stddev_pop_order_by represents the goals_stddev_pop_order_by GraphQL type
type goals_stddev_pop_order_by struct {
}

// goals_stddev_samp_order_by represents the goals_stddev_samp_order_by GraphQL type
type goals_stddev_samp_order_by struct {
}

// goals_stream_cursor_input represents the goals_stream_cursor_input GraphQL type
type goals_stream_cursor_input struct {
}

// goals_stream_cursor_value_input represents the goals_stream_cursor_value_input GraphQL type
type goals_stream_cursor_value_input struct {
}

// goals_sum_order_by represents the goals_sum_order_by GraphQL type
type goals_sum_order_by struct {
}

// goals_var_pop_order_by represents the goals_var_pop_order_by GraphQL type
type goals_var_pop_order_by struct {
}

// goals_var_samp_order_by represents the goals_var_samp_order_by GraphQL type
type goals_var_samp_order_by struct {
}

// goals_variance_order_by represents the goals_variance_order_by GraphQL type
type goals_variance_order_by struct {
}

// images represents the images GraphQL type
type images struct {
	Color string `json:"color"`
	Colors *json.RawMessage `json:"colors"`
	Height int `json:"height"`
	ID *bigint `json:"id"`
	Imageable_id int `json:"imageable_id"`
	Imageable_type string `json:"imageable_type"`
	Ratio *float8 `json:"ratio"`
	URL string `json:"url"`
	Width int `json:"width"`
}

// images_aggregate_order_by represents the images_aggregate_order_by GraphQL type
type images_aggregate_order_by struct {
}

// images_avg_order_by represents the images_avg_order_by GraphQL type
type images_avg_order_by struct {
}

// images_bool_exp represents the images_bool_exp GraphQL type
type images_bool_exp struct {
}

// images_max_order_by represents the images_max_order_by GraphQL type
type images_max_order_by struct {
}

// images_min_order_by represents the images_min_order_by GraphQL type
type images_min_order_by struct {
}

// images_order_by represents the images_order_by GraphQL type
type images_order_by struct {
}

// images_select_column represents the images_select_column GraphQL type
type images_select_column struct {
}

// images_stddev_order_by represents the images_stddev_order_by GraphQL type
type images_stddev_order_by struct {
}

// images_stddev_pop_order_by represents the images_stddev_pop_order_by GraphQL type
type images_stddev_pop_order_by struct {
}

// images_stddev_samp_order_by represents the images_stddev_samp_order_by GraphQL type
type images_stddev_samp_order_by struct {
}

// images_stream_cursor_input represents the images_stream_cursor_input GraphQL type
type images_stream_cursor_input struct {
}

// images_stream_cursor_value_input represents the images_stream_cursor_value_input GraphQL type
type images_stream_cursor_value_input struct {
}

// images_sum_order_by represents the images_sum_order_by GraphQL type
type images_sum_order_by struct {
}

// images_var_pop_order_by represents the images_var_pop_order_by GraphQL type
type images_var_pop_order_by struct {
}

// images_var_samp_order_by represents the images_var_samp_order_by GraphQL type
type images_var_samp_order_by struct {
}

// images_variance_order_by represents the images_variance_order_by GraphQL type
type images_variance_order_by struct {
}

// json_comparison_exp represents the json_comparison_exp GraphQL type
type json_comparison_exp struct {
}

// jsonb represents the jsonb GraphQL type
type jsonb struct {
}

// jsonb_cast_exp represents the jsonb_cast_exp GraphQL type
type jsonb_cast_exp struct {
}

// jsonb_comparison_exp represents the jsonb_comparison_exp GraphQL type
type jsonb_comparison_exp struct {
}

// languages represents the languages GraphQL type
type languages struct {
	Code2 string `json:"code2"`
	Code3 string `json:"code3"`
	ID int `json:"id"`
	Language string `json:"language"`
}

// languages_bool_exp represents the languages_bool_exp GraphQL type
type languages_bool_exp struct {
}

// languages_order_by represents the languages_order_by GraphQL type
type languages_order_by struct {
}

// languages_select_column represents the languages_select_column GraphQL type
type languages_select_column struct {
}

// languages_stream_cursor_input represents the languages_stream_cursor_input GraphQL type
type languages_stream_cursor_input struct {
}

// languages_stream_cursor_value_input represents the languages_stream_cursor_value_input GraphQL type
type languages_stream_cursor_value_input struct {
}

// likes represents the likes GraphQL type
type likes struct {
	Activity *activities `json:"activity"`
	Created_at *timestamptz `json:"created_at"`
	Followers []*followed_users `json:"followers"`
	ID int `json:"id"`
	Likeable_id int `json:"likeable_id"`
	Likeable_type string `json:"likeable_type"`
	List *lists `json:"list"`
	User *users `json:"user"`
	User_book *user_books `json:"user_book"`
	User_id int `json:"user_id"`
}

// likes_aggregate_order_by represents the likes_aggregate_order_by GraphQL type
type likes_aggregate_order_by struct {
}

// likes_avg_order_by represents the likes_avg_order_by GraphQL type
type likes_avg_order_by struct {
}

// likes_bool_exp represents the likes_bool_exp GraphQL type
type likes_bool_exp struct {
}

// likes_max_order_by represents the likes_max_order_by GraphQL type
type likes_max_order_by struct {
}

// likes_min_order_by represents the likes_min_order_by GraphQL type
type likes_min_order_by struct {
}

// likes_order_by represents the likes_order_by GraphQL type
type likes_order_by struct {
}

// likes_select_column represents the likes_select_column GraphQL type
type likes_select_column struct {
}

// likes_stddev_order_by represents the likes_stddev_order_by GraphQL type
type likes_stddev_order_by struct {
}

// likes_stddev_pop_order_by represents the likes_stddev_pop_order_by GraphQL type
type likes_stddev_pop_order_by struct {
}

// likes_stddev_samp_order_by represents the likes_stddev_samp_order_by GraphQL type
type likes_stddev_samp_order_by struct {
}

// likes_stream_cursor_input represents the likes_stream_cursor_input GraphQL type
type likes_stream_cursor_input struct {
}

// likes_stream_cursor_value_input represents the likes_stream_cursor_value_input GraphQL type
type likes_stream_cursor_value_input struct {
}

// likes_sum_order_by represents the likes_sum_order_by GraphQL type
type likes_sum_order_by struct {
}

// likes_var_pop_order_by represents the likes_var_pop_order_by GraphQL type
type likes_var_pop_order_by struct {
}

// likes_var_samp_order_by represents the likes_var_samp_order_by GraphQL type
type likes_var_samp_order_by struct {
}

// likes_variance_order_by represents the likes_variance_order_by GraphQL type
type likes_variance_order_by struct {
}

// list_books represents the list_books GraphQL type
type list_books struct {
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Created_at *timestamp `json:"created_at"`
	Date_added *timestamptz `json:"date_added"`
	Edition *editions `json:"edition"`
	Edition_id int `json:"edition_id"`
	ID int `json:"id"`
	Imported bool `json:"imported"`
	List *lists `json:"list"`
	List_id int `json:"list_id"`
	Merged_at *timestamp `json:"merged_at"`
	Original_book_id int `json:"original_book_id"`
	Original_edition_id int `json:"original_edition_id"`
	Position int `json:"position"`
	Reason string `json:"reason"`
	Updated_at *timestamptz `json:"updated_at"`
	User_books []*user_books `json:"user_books"`
	User_books_aggregate *user_books_aggregate `json:"user_books_aggregate"`
}

// list_books_aggregate represents the list_books_aggregate GraphQL type
type list_books_aggregate struct {
	Aggregate *list_books_aggregate_fields `json:"aggregate"`
	Nodes []*list_books `json:"nodes"`
}

// list_books_aggregate_bool_exp represents the list_books_aggregate_bool_exp GraphQL type
type list_books_aggregate_bool_exp struct {
}

// list_books_aggregate_bool_exp_bool_and represents the list_books_aggregate_bool_exp_bool_and GraphQL type
type list_books_aggregate_bool_exp_bool_and struct {
}

// list_books_aggregate_bool_exp_bool_or represents the list_books_aggregate_bool_exp_bool_or GraphQL type
type list_books_aggregate_bool_exp_bool_or struct {
}

// list_books_aggregate_bool_exp_count represents the list_books_aggregate_bool_exp_count GraphQL type
type list_books_aggregate_bool_exp_count struct {
}

// list_books_aggregate_fields represents the list_books_aggregate_fields GraphQL type
type list_books_aggregate_fields struct {
	Avg *list_books_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *list_books_max_fields `json:"max"`
	Min *list_books_min_fields `json:"min"`
	Stddev *list_books_stddev_fields `json:"stddev"`
	Stddev_pop *list_books_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *list_books_stddev_samp_fields `json:"stddev_samp"`
	Sum *list_books_sum_fields `json:"sum"`
	Var_pop *list_books_var_pop_fields `json:"var_pop"`
	Var_samp *list_books_var_samp_fields `json:"var_samp"`
	Variance *list_books_variance_fields `json:"variance"`
}

// list_books_aggregate_order_by represents the list_books_aggregate_order_by GraphQL type
type list_books_aggregate_order_by struct {
}

// list_books_avg_fields represents the list_books_avg_fields GraphQL type
type list_books_avg_fields struct {
	Book_id float64 `json:"book_id"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	List_id float64 `json:"list_id"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Position float64 `json:"position"`
}

// list_books_avg_order_by represents the list_books_avg_order_by GraphQL type
type list_books_avg_order_by struct {
}

// list_books_bool_exp represents the list_books_bool_exp GraphQL type
type list_books_bool_exp struct {
}

// list_books_inc_input represents the list_books_inc_input GraphQL type
type list_books_inc_input struct {
}

// list_books_max_fields represents the list_books_max_fields GraphQL type
type list_books_max_fields struct {
	Book_id int `json:"book_id"`
	Created_at *timestamp `json:"created_at"`
	Date_added *timestamptz `json:"date_added"`
	Edition_id int `json:"edition_id"`
	ID int `json:"id"`
	List_id int `json:"list_id"`
	Merged_at *timestamp `json:"merged_at"`
	Original_book_id int `json:"original_book_id"`
	Original_edition_id int `json:"original_edition_id"`
	Position int `json:"position"`
	Reason string `json:"reason"`
	Updated_at *timestamptz `json:"updated_at"`
}

// list_books_max_order_by represents the list_books_max_order_by GraphQL type
type list_books_max_order_by struct {
}

// list_books_min_fields represents the list_books_min_fields GraphQL type
type list_books_min_fields struct {
	Book_id int `json:"book_id"`
	Created_at *timestamp `json:"created_at"`
	Date_added *timestamptz `json:"date_added"`
	Edition_id int `json:"edition_id"`
	ID int `json:"id"`
	List_id int `json:"list_id"`
	Merged_at *timestamp `json:"merged_at"`
	Original_book_id int `json:"original_book_id"`
	Original_edition_id int `json:"original_edition_id"`
	Position int `json:"position"`
	Reason string `json:"reason"`
	Updated_at *timestamptz `json:"updated_at"`
}

// list_books_min_order_by represents the list_books_min_order_by GraphQL type
type list_books_min_order_by struct {
}

// list_books_mutation_response represents the list_books_mutation_response GraphQL type
type list_books_mutation_response struct {
	Affected_rows int `json:"affected_rows"`
	Returning []*list_books `json:"returning"`
}

// list_books_order_by represents the list_books_order_by GraphQL type
type list_books_order_by struct {
}

// list_books_pk_columns_input represents the list_books_pk_columns_input GraphQL type
type list_books_pk_columns_input struct {
}

// list_books_select_column represents the list_books_select_column GraphQL type
type list_books_select_column struct {
}

// list_books_select_column_list_books_aggregate_bool_exp_bool_and_arguments_columns represents the list_books_select_column_list_books_aggregate_bool_exp_bool_and_arguments_columns GraphQL type
type list_books_select_column_list_books_aggregate_bool_exp_bool_and_arguments_columns struct {
}

// list_books_select_column_list_books_aggregate_bool_exp_bool_or_arguments_columns represents the list_books_select_column_list_books_aggregate_bool_exp_bool_or_arguments_columns GraphQL type
type list_books_select_column_list_books_aggregate_bool_exp_bool_or_arguments_columns struct {
}

// list_books_set_input represents the list_books_set_input GraphQL type
type list_books_set_input struct {
}

// list_books_stddev_fields represents the list_books_stddev_fields GraphQL type
type list_books_stddev_fields struct {
	Book_id float64 `json:"book_id"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	List_id float64 `json:"list_id"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Position float64 `json:"position"`
}

// list_books_stddev_order_by represents the list_books_stddev_order_by GraphQL type
type list_books_stddev_order_by struct {
}

// list_books_stddev_pop_fields represents the list_books_stddev_pop_fields GraphQL type
type list_books_stddev_pop_fields struct {
	Book_id float64 `json:"book_id"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	List_id float64 `json:"list_id"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Position float64 `json:"position"`
}

// list_books_stddev_pop_order_by represents the list_books_stddev_pop_order_by GraphQL type
type list_books_stddev_pop_order_by struct {
}

// list_books_stddev_samp_fields represents the list_books_stddev_samp_fields GraphQL type
type list_books_stddev_samp_fields struct {
	Book_id float64 `json:"book_id"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	List_id float64 `json:"list_id"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Position float64 `json:"position"`
}

// list_books_stddev_samp_order_by represents the list_books_stddev_samp_order_by GraphQL type
type list_books_stddev_samp_order_by struct {
}

// list_books_stream_cursor_input represents the list_books_stream_cursor_input GraphQL type
type list_books_stream_cursor_input struct {
}

// list_books_stream_cursor_value_input represents the list_books_stream_cursor_value_input GraphQL type
type list_books_stream_cursor_value_input struct {
}

// list_books_sum_fields represents the list_books_sum_fields GraphQL type
type list_books_sum_fields struct {
	Book_id int `json:"book_id"`
	Edition_id int `json:"edition_id"`
	ID int `json:"id"`
	List_id int `json:"list_id"`
	Original_book_id int `json:"original_book_id"`
	Original_edition_id int `json:"original_edition_id"`
	Position int `json:"position"`
}

// list_books_sum_order_by represents the list_books_sum_order_by GraphQL type
type list_books_sum_order_by struct {
}

// list_books_updates represents the list_books_updates GraphQL type
type list_books_updates struct {
}

// list_books_var_pop_fields represents the list_books_var_pop_fields GraphQL type
type list_books_var_pop_fields struct {
	Book_id float64 `json:"book_id"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	List_id float64 `json:"list_id"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Position float64 `json:"position"`
}

// list_books_var_pop_order_by represents the list_books_var_pop_order_by GraphQL type
type list_books_var_pop_order_by struct {
}

// list_books_var_samp_fields represents the list_books_var_samp_fields GraphQL type
type list_books_var_samp_fields struct {
	Book_id float64 `json:"book_id"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	List_id float64 `json:"list_id"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Position float64 `json:"position"`
}

// list_books_var_samp_order_by represents the list_books_var_samp_order_by GraphQL type
type list_books_var_samp_order_by struct {
}

// list_books_variance_fields represents the list_books_variance_fields GraphQL type
type list_books_variance_fields struct {
	Book_id float64 `json:"book_id"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	List_id float64 `json:"list_id"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Position float64 `json:"position"`
}

// list_books_variance_order_by represents the list_books_variance_order_by GraphQL type
type list_books_variance_order_by struct {
}

// lists represents the lists GraphQL type
type lists struct {
	Books_count int `json:"books_count"`
	Created_at *timestamp `json:"created_at"`
	Default_view string `json:"default_view"`
	Description string `json:"description"`
	Featured bool `json:"featured"`
	Featured_profile bool `json:"featured_profile"`
	Followed_lists []*followed_lists `json:"followed_lists"`
	Followers []*followed_users `json:"followers"`
	Followers_count int `json:"followers_count"`
	ID int `json:"id"`
	Imported bool `json:"imported"`
	Likes []*likes `json:"likes"`
	Likes_count int `json:"likes_count"`
	List_books []*list_books `json:"list_books"`
	List_books_aggregate *list_books_aggregate `json:"list_books_aggregate"`
	Name string `json:"name"`
	Object_type string `json:"object_type"`
	Privacy_setting *privacy_settings `json:"privacy_setting"`
	Privacy_setting_id int `json:"privacy_setting_id"`
	Public bool `json:"public"`
	Ranked bool `json:"ranked"`
	Slug string `json:"slug"`
	Updated_at *timestamptz `json:"updated_at"`
	URL string `json:"url"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// lists_aggregate represents the lists_aggregate GraphQL type
type lists_aggregate struct {
	Aggregate *lists_aggregate_fields `json:"aggregate"`
	Nodes []*lists `json:"nodes"`
}

// lists_aggregate_bool_exp represents the lists_aggregate_bool_exp GraphQL type
type lists_aggregate_bool_exp struct {
}

// lists_aggregate_bool_exp_bool_and represents the lists_aggregate_bool_exp_bool_and GraphQL type
type lists_aggregate_bool_exp_bool_and struct {
}

// lists_aggregate_bool_exp_bool_or represents the lists_aggregate_bool_exp_bool_or GraphQL type
type lists_aggregate_bool_exp_bool_or struct {
}

// lists_aggregate_bool_exp_count represents the lists_aggregate_bool_exp_count GraphQL type
type lists_aggregate_bool_exp_count struct {
}

// lists_aggregate_fields represents the lists_aggregate_fields GraphQL type
type lists_aggregate_fields struct {
	Avg *lists_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *lists_max_fields `json:"max"`
	Min *lists_min_fields `json:"min"`
	Stddev *lists_stddev_fields `json:"stddev"`
	Stddev_pop *lists_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *lists_stddev_samp_fields `json:"stddev_samp"`
	Sum *lists_sum_fields `json:"sum"`
	Var_pop *lists_var_pop_fields `json:"var_pop"`
	Var_samp *lists_var_samp_fields `json:"var_samp"`
	Variance *lists_variance_fields `json:"variance"`
}

// lists_aggregate_order_by represents the lists_aggregate_order_by GraphQL type
type lists_aggregate_order_by struct {
}

// lists_avg_fields represents the lists_avg_fields GraphQL type
type lists_avg_fields struct {
	Books_count float64 `json:"books_count"`
	Followers_count float64 `json:"followers_count"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	User_id float64 `json:"user_id"`
}

// lists_avg_order_by represents the lists_avg_order_by GraphQL type
type lists_avg_order_by struct {
}

// lists_bool_exp represents the lists_bool_exp GraphQL type
type lists_bool_exp struct {
}

// lists_max_fields represents the lists_max_fields GraphQL type
type lists_max_fields struct {
	Books_count int `json:"books_count"`
	Created_at *timestamp `json:"created_at"`
	Default_view string `json:"default_view"`
	Description string `json:"description"`
	Followers_count int `json:"followers_count"`
	ID int `json:"id"`
	Likes_count int `json:"likes_count"`
	Name string `json:"name"`
	Object_type string `json:"object_type"`
	Privacy_setting_id int `json:"privacy_setting_id"`
	Slug string `json:"slug"`
	Updated_at *timestamptz `json:"updated_at"`
	URL string `json:"url"`
	User_id int `json:"user_id"`
}

// lists_max_order_by represents the lists_max_order_by GraphQL type
type lists_max_order_by struct {
}

// lists_min_fields represents the lists_min_fields GraphQL type
type lists_min_fields struct {
	Books_count int `json:"books_count"`
	Created_at *timestamp `json:"created_at"`
	Default_view string `json:"default_view"`
	Description string `json:"description"`
	Followers_count int `json:"followers_count"`
	ID int `json:"id"`
	Likes_count int `json:"likes_count"`
	Name string `json:"name"`
	Object_type string `json:"object_type"`
	Privacy_setting_id int `json:"privacy_setting_id"`
	Slug string `json:"slug"`
	Updated_at *timestamptz `json:"updated_at"`
	URL string `json:"url"`
	User_id int `json:"user_id"`
}

// lists_min_order_by represents the lists_min_order_by GraphQL type
type lists_min_order_by struct {
}

// lists_order_by represents the lists_order_by GraphQL type
type lists_order_by struct {
}

// lists_select_column represents the lists_select_column GraphQL type
type lists_select_column struct {
}

// lists_select_column_lists_aggregate_bool_exp_bool_and_arguments_columns represents the lists_select_column_lists_aggregate_bool_exp_bool_and_arguments_columns GraphQL type
type lists_select_column_lists_aggregate_bool_exp_bool_and_arguments_columns struct {
}

// lists_select_column_lists_aggregate_bool_exp_bool_or_arguments_columns represents the lists_select_column_lists_aggregate_bool_exp_bool_or_arguments_columns GraphQL type
type lists_select_column_lists_aggregate_bool_exp_bool_or_arguments_columns struct {
}

// lists_stddev_fields represents the lists_stddev_fields GraphQL type
type lists_stddev_fields struct {
	Books_count float64 `json:"books_count"`
	Followers_count float64 `json:"followers_count"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	User_id float64 `json:"user_id"`
}

// lists_stddev_order_by represents the lists_stddev_order_by GraphQL type
type lists_stddev_order_by struct {
}

// lists_stddev_pop_fields represents the lists_stddev_pop_fields GraphQL type
type lists_stddev_pop_fields struct {
	Books_count float64 `json:"books_count"`
	Followers_count float64 `json:"followers_count"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	User_id float64 `json:"user_id"`
}

// lists_stddev_pop_order_by represents the lists_stddev_pop_order_by GraphQL type
type lists_stddev_pop_order_by struct {
}

// lists_stddev_samp_fields represents the lists_stddev_samp_fields GraphQL type
type lists_stddev_samp_fields struct {
	Books_count float64 `json:"books_count"`
	Followers_count float64 `json:"followers_count"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	User_id float64 `json:"user_id"`
}

// lists_stddev_samp_order_by represents the lists_stddev_samp_order_by GraphQL type
type lists_stddev_samp_order_by struct {
}

// lists_stream_cursor_input represents the lists_stream_cursor_input GraphQL type
type lists_stream_cursor_input struct {
}

// lists_stream_cursor_value_input represents the lists_stream_cursor_value_input GraphQL type
type lists_stream_cursor_value_input struct {
}

// lists_sum_fields represents the lists_sum_fields GraphQL type
type lists_sum_fields struct {
	Books_count int `json:"books_count"`
	Followers_count int `json:"followers_count"`
	ID int `json:"id"`
	Likes_count int `json:"likes_count"`
	Privacy_setting_id int `json:"privacy_setting_id"`
	User_id int `json:"user_id"`
}

// lists_sum_order_by represents the lists_sum_order_by GraphQL type
type lists_sum_order_by struct {
}

// lists_var_pop_fields represents the lists_var_pop_fields GraphQL type
type lists_var_pop_fields struct {
	Books_count float64 `json:"books_count"`
	Followers_count float64 `json:"followers_count"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	User_id float64 `json:"user_id"`
}

// lists_var_pop_order_by represents the lists_var_pop_order_by GraphQL type
type lists_var_pop_order_by struct {
}

// lists_var_samp_fields represents the lists_var_samp_fields GraphQL type
type lists_var_samp_fields struct {
	Books_count float64 `json:"books_count"`
	Followers_count float64 `json:"followers_count"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	User_id float64 `json:"user_id"`
}

// lists_var_samp_order_by represents the lists_var_samp_order_by GraphQL type
type lists_var_samp_order_by struct {
}

// lists_variance_fields represents the lists_variance_fields GraphQL type
type lists_variance_fields struct {
	Books_count float64 `json:"books_count"`
	Followers_count float64 `json:"followers_count"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	User_id float64 `json:"user_id"`
}

// lists_variance_order_by represents the lists_variance_order_by GraphQL type
type lists_variance_order_by struct {
}

// mutation_root represents the mutation_root GraphQL type
type mutation_root struct {
	Book_mapping_normalize *BookMappingIdType `json:"book_mapping_normalize"`
	Book_normalize *BookIdType `json:"book_normalize"`
	Collection_import_result_reimport *CollectionImportResultIdType `json:"collection_import_result_reimport"`
	Collection_import_retry *CollectionImportIdType `json:"collection_import_retry"`
	Delete_activities *activities_mutation_response `json:"delete_activities"`
	Delete_activities_by_pk *activities `json:"delete_activities_by_pk"`
	Delete_book_mapping *BookMappingIdType `json:"delete_book_mapping"`
	Delete_followed_list *DeleteListType `json:"delete_followed_list"`
	Delete_followed_prompt *DeleteFollowedPromptType `json:"delete_followed_prompt"`
	Delete_followed_prompts *followed_prompts_mutation_response `json:"delete_followed_prompts"`
	Delete_followed_prompts_by_pk *followed_prompts `json:"delete_followed_prompts_by_pk"`
	Delete_followed_user *FollowedUserType `json:"delete_followed_user"`
	Delete_followed_users *followed_users_mutation_response `json:"delete_followed_users"`
	Delete_followed_users_by_pk *followed_users `json:"delete_followed_users_by_pk"`
	Delete_goals *goals_mutation_response `json:"delete_goals"`
	Delete_goals_by_pk *goals `json:"delete_goals_by_pk"`
	Delete_like *LikeDeleteType `json:"delete_like"`
	Delete_list *ListDeleteType `json:"delete_list"`
	Delete_list_book *ListBookDeleteType `json:"delete_list_book"`
	Delete_prompt_answer *PromptAnswerIdType `json:"delete_prompt_answer"`
	Delete_prompts *prompts_mutation_response `json:"delete_prompts"`
	Delete_prompts_by_pk *prompts `json:"delete_prompts_by_pk"`
	Delete_reading_journal *DeleteReadingJournalOutput `json:"delete_reading_journal"`
	Delete_reading_journals_for_book *DeleteReadingJournalsOutput `json:"delete_reading_journals_for_book"`
	Delete_user_blocks *user_blocks_mutation_response `json:"delete_user_blocks"`
	Delete_user_blocks_by_pk *user_blocks `json:"delete_user_blocks_by_pk"`
	Delete_user_book *UserBookDeleteType `json:"delete_user_book"`
	Delete_user_book_read *UserBookReadIdType `json:"delete_user_book_read"`
	Edition_normalize *EditionIdType `json:"edition_normalize"`
	Edition_owned *ListBookIdType `json:"edition_owned"`
	Email_user_delete_confirmation *SuccessType `json:"email_user_delete_confirmation"`
	Insert_author *AuthorIdType `json:"insert_author"`
	Insert_block *InsertBlockOutput `json:"insert_block"`
	Insert_book *OptionalEditionIdType `json:"insert_book"`
	Insert_book_mapping *BookMappingIdType `json:"insert_book_mapping"`
	Insert_character *CharacterIdType `json:"insert_character"`
	Insert_collection_import *CollectionImportIdType `json:"insert_collection_import"`
	Insert_edition *EditionIdType `json:"insert_edition"`
	Insert_followed_prompts *followed_prompts_mutation_response `json:"insert_followed_prompts"`
	Insert_followed_prompts_one *followed_prompts `json:"insert_followed_prompts_one"`
	Insert_followed_user *FollowedUserType `json:"insert_followed_user"`
	Insert_goal *GoalIdType `json:"insert_goal"`
	Insert_image *ImageIdType `json:"insert_image"`
	Insert_list *ListIdType `json:"insert_list"`
	Insert_list_book *ListBookIdType `json:"insert_list_book"`
	Insert_notification_settings *notification_settings_mutation_response `json:"insert_notification_settings"`
	Insert_notification_settings_one *notification_settings `json:"insert_notification_settings_one"`
	Insert_prompt *PromptIdType `json:"insert_prompt"`
	Insert_prompt_answer *PromptAnswerIdType `json:"insert_prompt_answer"`
	Insert_publisher *PublisherIdType `json:"insert_publisher"`
	Insert_reading_journal *ReadingJournalOutput `json:"insert_reading_journal"`
	Insert_report *ReportOutput `json:"insert_report"`
	Insert_serie *SeriesIdType `json:"insert_serie"`
	Insert_user *UserIdType `json:"insert_user"`
	Insert_user_blocks *user_blocks_mutation_response `json:"insert_user_blocks"`
	Insert_user_blocks_one *user_blocks `json:"insert_user_blocks_one"`
	Insert_user_book *UserBookIdType `json:"insert_user_book"`
	Insert_user_book_read *UserBookReadIdType `json:"insert_user_book_read"`
	Insert_user_flags *user_flags_mutation_response `json:"insert_user_flags"`
	Insert_user_flags_one *user_flags `json:"insert_user_flags_one"`
	Receipt_validate *ValidateReceiptType `json:"receipt_validate"`
	Update_author *AuthorIdType `json:"update_author"`
	Update_book *BookIdType `json:"update_book"`
	Update_character *CharacterIdType `json:"update_character"`
	Update_collection_import_results *collection_import_results_mutation_response `json:"update_collection_import_results"`
	Update_collection_import_results_by_pk *collection_import_results `json:"update_collection_import_results_by_pk"`
	Update_collection_import_results_many []*collection_import_results_mutation_response `json:"update_collection_import_results_many"`
	Update_edition *EditionIdType `json:"update_edition"`
	Update_followed_prompts *followed_prompts_mutation_response `json:"update_followed_prompts"`
	Update_followed_prompts_by_pk *followed_prompts `json:"update_followed_prompts_by_pk"`
	Update_followed_prompts_many []*followed_prompts_mutation_response `json:"update_followed_prompts_many"`
	Update_goal *GoalIdType `json:"update_goal"`
	Update_goal_progress *GoalIdType `json:"update_goal_progress"`
	Update_list *ListIdType `json:"update_list"`
	Update_list_books *list_books_mutation_response `json:"update_list_books"`
	Update_list_books_by_pk *list_books `json:"update_list_books_by_pk"`
	Update_list_books_many []*list_books_mutation_response `json:"update_list_books_many"`
	Update_newsletter *NewsletterStatusType `json:"update_newsletter"`
	Update_notification_deliveries *notification_deliveries_mutation_response `json:"update_notification_deliveries"`
	Update_notification_deliveries_by_pk *notification_deliveries `json:"update_notification_deliveries_by_pk"`
	Update_notification_deliveries_many []*notification_deliveries_mutation_response `json:"update_notification_deliveries_many"`
	Update_notification_settings *notification_settings_mutation_response `json:"update_notification_settings"`
	Update_notification_settings_by_pk *notification_settings `json:"update_notification_settings_by_pk"`
	Update_notification_settings_many []*notification_settings_mutation_response `json:"update_notification_settings_many"`
	Update_prompt *PromptIdType `json:"update_prompt"`
	Update_prompt_answers *prompt_answers_mutation_response `json:"update_prompt_answers"`
	Update_prompt_answers_by_pk *prompt_answers `json:"update_prompt_answers_by_pk"`
	Update_prompt_answers_many []*prompt_answers_mutation_response `json:"update_prompt_answers_many"`
	Update_publisher *PublisherIdType `json:"update_publisher"`
	Update_reading_journal *ReadingJournalOutput `json:"update_reading_journal"`
	Update_serie *SeriesIdType `json:"update_serie"`
	Update_user *UserIdType `json:"update_user"`
	Update_user_book *UserBookIdType `json:"update_user_book"`
	Update_user_book_read *UserBookReadIdType `json:"update_user_book_read"`
	Update_user_privacy_setting *UserIdType `json:"update_user_privacy_setting"`
	Upsert_book *NewBookIdType `json:"upsert_book"`
	Upsert_followed_list *FollowedListType `json:"upsert_followed_list"`
	Upsert_followed_prompt *FollowedPromptType `json:"upsert_followed_prompt"`
	Upsert_like *LikeType `json:"upsert_like"`
	Upsert_tags *TagsType `json:"upsert_tags"`
	Upsert_user_book_reads *UserBooksReadUpsertType `json:"upsert_user_book_reads"`
	User_login *UserIdType `json:"user_login"`
}

// notification_channels represents the notification_channels GraphQL type
type notification_channels struct {
	Channel string `json:"channel"`
	ID *bigint `json:"id"`
}

// notification_channels_bool_exp represents the notification_channels_bool_exp GraphQL type
type notification_channels_bool_exp struct {
}

// notification_channels_order_by represents the notification_channels_order_by GraphQL type
type notification_channels_order_by struct {
}

// notification_channels_select_column represents the notification_channels_select_column GraphQL type
type notification_channels_select_column struct {
}

// notification_channels_stream_cursor_input represents the notification_channels_stream_cursor_input GraphQL type
type notification_channels_stream_cursor_input struct {
}

// notification_channels_stream_cursor_value_input represents the notification_channels_stream_cursor_value_input GraphQL type
type notification_channels_stream_cursor_value_input struct {
}

// notification_deliveries represents the notification_deliveries GraphQL type
type notification_deliveries struct {
	Channel *notification_channels `json:"channel"`
	Channel_id int `json:"channel_id"`
	ID *bigint `json:"id"`
	Notification *notifications `json:"notification"`
	Notification_id int `json:"notification_id"`
	Read bool `json:"read"`
	Read_at *timestamp `json:"read_at"`
	Sent_at *timestamp `json:"sent_at"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// notification_deliveries_aggregate represents the notification_deliveries_aggregate GraphQL type
type notification_deliveries_aggregate struct {
	Aggregate *notification_deliveries_aggregate_fields `json:"aggregate"`
	Nodes []*notification_deliveries `json:"nodes"`
}

// notification_deliveries_aggregate_bool_exp represents the notification_deliveries_aggregate_bool_exp GraphQL type
type notification_deliveries_aggregate_bool_exp struct {
}

// notification_deliveries_aggregate_bool_exp_bool_and represents the notification_deliveries_aggregate_bool_exp_bool_and GraphQL type
type notification_deliveries_aggregate_bool_exp_bool_and struct {
}

// notification_deliveries_aggregate_bool_exp_bool_or represents the notification_deliveries_aggregate_bool_exp_bool_or GraphQL type
type notification_deliveries_aggregate_bool_exp_bool_or struct {
}

// notification_deliveries_aggregate_bool_exp_count represents the notification_deliveries_aggregate_bool_exp_count GraphQL type
type notification_deliveries_aggregate_bool_exp_count struct {
}

// notification_deliveries_aggregate_fields represents the notification_deliveries_aggregate_fields GraphQL type
type notification_deliveries_aggregate_fields struct {
	Avg *notification_deliveries_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *notification_deliveries_max_fields `json:"max"`
	Min *notification_deliveries_min_fields `json:"min"`
	Stddev *notification_deliveries_stddev_fields `json:"stddev"`
	Stddev_pop *notification_deliveries_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *notification_deliveries_stddev_samp_fields `json:"stddev_samp"`
	Sum *notification_deliveries_sum_fields `json:"sum"`
	Var_pop *notification_deliveries_var_pop_fields `json:"var_pop"`
	Var_samp *notification_deliveries_var_samp_fields `json:"var_samp"`
	Variance *notification_deliveries_variance_fields `json:"variance"`
}

// notification_deliveries_aggregate_order_by represents the notification_deliveries_aggregate_order_by GraphQL type
type notification_deliveries_aggregate_order_by struct {
}

// notification_deliveries_avg_fields represents the notification_deliveries_avg_fields GraphQL type
type notification_deliveries_avg_fields struct {
	Channel_id float64 `json:"channel_id"`
	ID float64 `json:"id"`
	Notification_id float64 `json:"notification_id"`
	User_id float64 `json:"user_id"`
}

// notification_deliveries_avg_order_by represents the notification_deliveries_avg_order_by GraphQL type
type notification_deliveries_avg_order_by struct {
}

// notification_deliveries_bool_exp represents the notification_deliveries_bool_exp GraphQL type
type notification_deliveries_bool_exp struct {
}

// notification_deliveries_max_fields represents the notification_deliveries_max_fields GraphQL type
type notification_deliveries_max_fields struct {
	Channel_id int `json:"channel_id"`
	ID *bigint `json:"id"`
	Notification_id int `json:"notification_id"`
	Read_at *timestamp `json:"read_at"`
	Sent_at *timestamp `json:"sent_at"`
	User_id int `json:"user_id"`
}

// notification_deliveries_max_order_by represents the notification_deliveries_max_order_by GraphQL type
type notification_deliveries_max_order_by struct {
}

// notification_deliveries_min_fields represents the notification_deliveries_min_fields GraphQL type
type notification_deliveries_min_fields struct {
	Channel_id int `json:"channel_id"`
	ID *bigint `json:"id"`
	Notification_id int `json:"notification_id"`
	Read_at *timestamp `json:"read_at"`
	Sent_at *timestamp `json:"sent_at"`
	User_id int `json:"user_id"`
}

// notification_deliveries_min_order_by represents the notification_deliveries_min_order_by GraphQL type
type notification_deliveries_min_order_by struct {
}

// notification_deliveries_mutation_response represents the notification_deliveries_mutation_response GraphQL type
type notification_deliveries_mutation_response struct {
	Affected_rows int `json:"affected_rows"`
	Returning []*notification_deliveries `json:"returning"`
}

// notification_deliveries_order_by represents the notification_deliveries_order_by GraphQL type
type notification_deliveries_order_by struct {
}

// notification_deliveries_pk_columns_input represents the notification_deliveries_pk_columns_input GraphQL type
type notification_deliveries_pk_columns_input struct {
}

// notification_deliveries_select_column represents the notification_deliveries_select_column GraphQL type
type notification_deliveries_select_column struct {
}

// notification_deliveries_select_column_notification_deliveries_aggregate_bool_exp_bool_and_arguments_columns represents the notification_deliveries_select_column_notification_deliveries_aggregate_bool_exp_bool_and_arguments_columns GraphQL type
type notification_deliveries_select_column_notification_deliveries_aggregate_bool_exp_bool_and_arguments_columns struct {
}

// notification_deliveries_select_column_notification_deliveries_aggregate_bool_exp_bool_or_arguments_columns represents the notification_deliveries_select_column_notification_deliveries_aggregate_bool_exp_bool_or_arguments_columns GraphQL type
type notification_deliveries_select_column_notification_deliveries_aggregate_bool_exp_bool_or_arguments_columns struct {
}

// notification_deliveries_set_input represents the notification_deliveries_set_input GraphQL type
type notification_deliveries_set_input struct {
}

// notification_deliveries_stddev_fields represents the notification_deliveries_stddev_fields GraphQL type
type notification_deliveries_stddev_fields struct {
	Channel_id float64 `json:"channel_id"`
	ID float64 `json:"id"`
	Notification_id float64 `json:"notification_id"`
	User_id float64 `json:"user_id"`
}

// notification_deliveries_stddev_order_by represents the notification_deliveries_stddev_order_by GraphQL type
type notification_deliveries_stddev_order_by struct {
}

// notification_deliveries_stddev_pop_fields represents the notification_deliveries_stddev_pop_fields GraphQL type
type notification_deliveries_stddev_pop_fields struct {
	Channel_id float64 `json:"channel_id"`
	ID float64 `json:"id"`
	Notification_id float64 `json:"notification_id"`
	User_id float64 `json:"user_id"`
}

// notification_deliveries_stddev_pop_order_by represents the notification_deliveries_stddev_pop_order_by GraphQL type
type notification_deliveries_stddev_pop_order_by struct {
}

// notification_deliveries_stddev_samp_fields represents the notification_deliveries_stddev_samp_fields GraphQL type
type notification_deliveries_stddev_samp_fields struct {
	Channel_id float64 `json:"channel_id"`
	ID float64 `json:"id"`
	Notification_id float64 `json:"notification_id"`
	User_id float64 `json:"user_id"`
}

// notification_deliveries_stddev_samp_order_by represents the notification_deliveries_stddev_samp_order_by GraphQL type
type notification_deliveries_stddev_samp_order_by struct {
}

// notification_deliveries_stream_cursor_input represents the notification_deliveries_stream_cursor_input GraphQL type
type notification_deliveries_stream_cursor_input struct {
}

// notification_deliveries_stream_cursor_value_input represents the notification_deliveries_stream_cursor_value_input GraphQL type
type notification_deliveries_stream_cursor_value_input struct {
}

// notification_deliveries_sum_fields represents the notification_deliveries_sum_fields GraphQL type
type notification_deliveries_sum_fields struct {
	Channel_id int `json:"channel_id"`
	ID *bigint `json:"id"`
	Notification_id int `json:"notification_id"`
	User_id int `json:"user_id"`
}

// notification_deliveries_sum_order_by represents the notification_deliveries_sum_order_by GraphQL type
type notification_deliveries_sum_order_by struct {
}

// notification_deliveries_updates represents the notification_deliveries_updates GraphQL type
type notification_deliveries_updates struct {
}

// notification_deliveries_var_pop_fields represents the notification_deliveries_var_pop_fields GraphQL type
type notification_deliveries_var_pop_fields struct {
	Channel_id float64 `json:"channel_id"`
	ID float64 `json:"id"`
	Notification_id float64 `json:"notification_id"`
	User_id float64 `json:"user_id"`
}

// notification_deliveries_var_pop_order_by represents the notification_deliveries_var_pop_order_by GraphQL type
type notification_deliveries_var_pop_order_by struct {
}

// notification_deliveries_var_samp_fields represents the notification_deliveries_var_samp_fields GraphQL type
type notification_deliveries_var_samp_fields struct {
	Channel_id float64 `json:"channel_id"`
	ID float64 `json:"id"`
	Notification_id float64 `json:"notification_id"`
	User_id float64 `json:"user_id"`
}

// notification_deliveries_var_samp_order_by represents the notification_deliveries_var_samp_order_by GraphQL type
type notification_deliveries_var_samp_order_by struct {
}

// notification_deliveries_variance_fields represents the notification_deliveries_variance_fields GraphQL type
type notification_deliveries_variance_fields struct {
	Channel_id float64 `json:"channel_id"`
	ID float64 `json:"id"`
	Notification_id float64 `json:"notification_id"`
	User_id float64 `json:"user_id"`
}

// notification_deliveries_variance_order_by represents the notification_deliveries_variance_order_by GraphQL type
type notification_deliveries_variance_order_by struct {
}

// notification_settings represents the notification_settings GraphQL type
type notification_settings struct {
	Channel_ids *json.RawMessage `json:"channel_ids"`
	ID *bigint `json:"id"`
	Notification_type_id int `json:"notification_type_id"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// notification_settings_aggregate_order_by represents the notification_settings_aggregate_order_by GraphQL type
type notification_settings_aggregate_order_by struct {
}

// notification_settings_avg_order_by represents the notification_settings_avg_order_by GraphQL type
type notification_settings_avg_order_by struct {
}

// notification_settings_bool_exp represents the notification_settings_bool_exp GraphQL type
type notification_settings_bool_exp struct {
}

// notification_settings_constraint represents the notification_settings_constraint GraphQL type
type notification_settings_constraint struct {
}

// notification_settings_insert_input represents the notification_settings_insert_input GraphQL type
type notification_settings_insert_input struct {
}

// notification_settings_max_order_by represents the notification_settings_max_order_by GraphQL type
type notification_settings_max_order_by struct {
}

// notification_settings_min_order_by represents the notification_settings_min_order_by GraphQL type
type notification_settings_min_order_by struct {
}

// notification_settings_mutation_response represents the notification_settings_mutation_response GraphQL type
type notification_settings_mutation_response struct {
	Affected_rows int `json:"affected_rows"`
	Returning []*notification_settings `json:"returning"`
}

// notification_settings_on_conflict represents the notification_settings_on_conflict GraphQL type
type notification_settings_on_conflict struct {
}

// notification_settings_order_by represents the notification_settings_order_by GraphQL type
type notification_settings_order_by struct {
}

// notification_settings_pk_columns_input represents the notification_settings_pk_columns_input GraphQL type
type notification_settings_pk_columns_input struct {
}

// notification_settings_select_column represents the notification_settings_select_column GraphQL type
type notification_settings_select_column struct {
}

// notification_settings_set_input represents the notification_settings_set_input GraphQL type
type notification_settings_set_input struct {
}

// notification_settings_stddev_order_by represents the notification_settings_stddev_order_by GraphQL type
type notification_settings_stddev_order_by struct {
}

// notification_settings_stddev_pop_order_by represents the notification_settings_stddev_pop_order_by GraphQL type
type notification_settings_stddev_pop_order_by struct {
}

// notification_settings_stddev_samp_order_by represents the notification_settings_stddev_samp_order_by GraphQL type
type notification_settings_stddev_samp_order_by struct {
}

// notification_settings_stream_cursor_input represents the notification_settings_stream_cursor_input GraphQL type
type notification_settings_stream_cursor_input struct {
}

// notification_settings_stream_cursor_value_input represents the notification_settings_stream_cursor_value_input GraphQL type
type notification_settings_stream_cursor_value_input struct {
}

// notification_settings_sum_order_by represents the notification_settings_sum_order_by GraphQL type
type notification_settings_sum_order_by struct {
}

// notification_settings_update_column represents the notification_settings_update_column GraphQL type
type notification_settings_update_column struct {
}

// notification_settings_updates represents the notification_settings_updates GraphQL type
type notification_settings_updates struct {
}

// notification_settings_var_pop_order_by represents the notification_settings_var_pop_order_by GraphQL type
type notification_settings_var_pop_order_by struct {
}

// notification_settings_var_samp_order_by represents the notification_settings_var_samp_order_by GraphQL type
type notification_settings_var_samp_order_by struct {
}

// notification_settings_variance_order_by represents the notification_settings_variance_order_by GraphQL type
type notification_settings_variance_order_by struct {
}

// notification_types represents the notification_types GraphQL type
type notification_types struct {
	Active bool `json:"active"`
	Default_channel_ids *json.RawMessage `json:"default_channel_ids"`
	Default_priority int `json:"default_priority"`
	Description string `json:"description"`
	ID *bigint `json:"id"`
	Name string `json:"name"`
	Notification_settings []*notification_settings `json:"notification_settings"`
	Uid string `json:"uid"`
}

// notification_types_bool_exp represents the notification_types_bool_exp GraphQL type
type notification_types_bool_exp struct {
}

// notification_types_order_by represents the notification_types_order_by GraphQL type
type notification_types_order_by struct {
}

// notification_types_select_column represents the notification_types_select_column GraphQL type
type notification_types_select_column struct {
}

// notification_types_stream_cursor_input represents the notification_types_stream_cursor_input GraphQL type
type notification_types_stream_cursor_input struct {
}

// notification_types_stream_cursor_value_input represents the notification_types_stream_cursor_value_input GraphQL type
type notification_types_stream_cursor_value_input struct {
}

// notifications represents the notifications GraphQL type
type notifications struct {
	Created_at *timestamptz `json:"created_at"`
	Description string `json:"description"`
	ID int `json:"id"`
	Link string `json:"link"`
	Link_text string `json:"link_text"`
	Notification_deliveries []*notification_deliveries `json:"notification_deliveries"`
	Notification_deliveries_aggregate *notification_deliveries_aggregate `json:"notification_deliveries_aggregate"`
	Notification_type_id int `json:"notification_type_id"`
	NotifierUser *users `json:"notifierUser"`
	Notifier_user_id int `json:"notifier_user_id"`
	Priority int `json:"priority"`
	Title string `json:"title"`
	Uid string `json:"uid"`
}

// notifications_bool_exp represents the notifications_bool_exp GraphQL type
type notifications_bool_exp struct {
}

// notifications_order_by represents the notifications_order_by GraphQL type
type notifications_order_by struct {
}

// notifications_select_column represents the notifications_select_column GraphQL type
type notifications_select_column struct {
}

// notifications_stream_cursor_input represents the notifications_stream_cursor_input GraphQL type
type notifications_stream_cursor_input struct {
}

// notifications_stream_cursor_value_input represents the notifications_stream_cursor_value_input GraphQL type
type notifications_stream_cursor_value_input struct {
}

// numeric represents the numeric GraphQL type
type numeric struct {
}

// numeric_comparison_exp represents the numeric_comparison_exp GraphQL type
type numeric_comparison_exp struct {
}

// order_by represents the order_by GraphQL type
type order_by struct {
}

// platforms represents the platforms GraphQL type
type platforms struct {
	Book_mappings []*book_mappings `json:"book_mappings"`
	ID int `json:"id"`
	Name string `json:"name"`
	URL string `json:"url"`
}

// platforms_bool_exp represents the platforms_bool_exp GraphQL type
type platforms_bool_exp struct {
}

// platforms_order_by represents the platforms_order_by GraphQL type
type platforms_order_by struct {
}

// platforms_select_column represents the platforms_select_column GraphQL type
type platforms_select_column struct {
}

// platforms_stream_cursor_input represents the platforms_stream_cursor_input GraphQL type
type platforms_stream_cursor_input struct {
}

// platforms_stream_cursor_value_input represents the platforms_stream_cursor_value_input GraphQL type
type platforms_stream_cursor_value_input struct {
}

// privacy_settings represents the privacy_settings GraphQL type
type privacy_settings struct {
	Activities []*activities `json:"activities"`
	ID int `json:"id"`
	Lists []*lists `json:"lists"`
	Lists_aggregate *lists_aggregate `json:"lists_aggregate"`
	Prompts []*prompts `json:"prompts"`
	Setting string `json:"setting"`
	User_books []*user_books `json:"user_books"`
	User_books_aggregate *user_books_aggregate `json:"user_books_aggregate"`
	Users []*users `json:"users"`
	Users_by_activity []*users `json:"users_by_activity"`
}

// privacy_settings_bool_exp represents the privacy_settings_bool_exp GraphQL type
type privacy_settings_bool_exp struct {
}

// privacy_settings_order_by represents the privacy_settings_order_by GraphQL type
type privacy_settings_order_by struct {
}

// privacy_settings_select_column represents the privacy_settings_select_column GraphQL type
type privacy_settings_select_column struct {
}

// privacy_settings_stream_cursor_input represents the privacy_settings_stream_cursor_input GraphQL type
type privacy_settings_stream_cursor_input struct {
}

// privacy_settings_stream_cursor_value_input represents the privacy_settings_stream_cursor_value_input GraphQL type
type privacy_settings_stream_cursor_value_input struct {
}

// prompt_answers represents the prompt_answers GraphQL type
type prompt_answers struct {
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Created_at *timestamptz `json:"created_at"`
	Description string `json:"description"`
	ID int `json:"id"`
	Merged_at *timestamp `json:"merged_at"`
	Original_book_id int `json:"original_book_id"`
	Prompt *prompts `json:"prompt"`
	Prompt_book *prompt_books_summary `json:"prompt_book"`
	Prompt_id int `json:"prompt_id"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// prompt_answers_aggregate represents the prompt_answers_aggregate GraphQL type
type prompt_answers_aggregate struct {
	Aggregate *prompt_answers_aggregate_fields `json:"aggregate"`
	Nodes []*prompt_answers `json:"nodes"`
}

// prompt_answers_aggregate_bool_exp represents the prompt_answers_aggregate_bool_exp GraphQL type
type prompt_answers_aggregate_bool_exp struct {
}

// prompt_answers_aggregate_bool_exp_count represents the prompt_answers_aggregate_bool_exp_count GraphQL type
type prompt_answers_aggregate_bool_exp_count struct {
}

// prompt_answers_aggregate_fields represents the prompt_answers_aggregate_fields GraphQL type
type prompt_answers_aggregate_fields struct {
	Avg *prompt_answers_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *prompt_answers_max_fields `json:"max"`
	Min *prompt_answers_min_fields `json:"min"`
	Stddev *prompt_answers_stddev_fields `json:"stddev"`
	Stddev_pop *prompt_answers_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *prompt_answers_stddev_samp_fields `json:"stddev_samp"`
	Sum *prompt_answers_sum_fields `json:"sum"`
	Var_pop *prompt_answers_var_pop_fields `json:"var_pop"`
	Var_samp *prompt_answers_var_samp_fields `json:"var_samp"`
	Variance *prompt_answers_variance_fields `json:"variance"`
}

// prompt_answers_aggregate_order_by represents the prompt_answers_aggregate_order_by GraphQL type
type prompt_answers_aggregate_order_by struct {
}

// prompt_answers_avg_fields represents the prompt_answers_avg_fields GraphQL type
type prompt_answers_avg_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Original_book_id float64 `json:"original_book_id"`
	Prompt_id float64 `json:"prompt_id"`
	User_id float64 `json:"user_id"`
}

// prompt_answers_avg_order_by represents the prompt_answers_avg_order_by GraphQL type
type prompt_answers_avg_order_by struct {
}

// prompt_answers_bool_exp represents the prompt_answers_bool_exp GraphQL type
type prompt_answers_bool_exp struct {
}

// prompt_answers_max_fields represents the prompt_answers_max_fields GraphQL type
type prompt_answers_max_fields struct {
	Book_id int `json:"book_id"`
	Created_at *timestamptz `json:"created_at"`
	Description string `json:"description"`
	ID int `json:"id"`
	Merged_at *timestamp `json:"merged_at"`
	Original_book_id int `json:"original_book_id"`
	Prompt_id int `json:"prompt_id"`
	User_id int `json:"user_id"`
}

// prompt_answers_max_order_by represents the prompt_answers_max_order_by GraphQL type
type prompt_answers_max_order_by struct {
}

// prompt_answers_min_fields represents the prompt_answers_min_fields GraphQL type
type prompt_answers_min_fields struct {
	Book_id int `json:"book_id"`
	Created_at *timestamptz `json:"created_at"`
	Description string `json:"description"`
	ID int `json:"id"`
	Merged_at *timestamp `json:"merged_at"`
	Original_book_id int `json:"original_book_id"`
	Prompt_id int `json:"prompt_id"`
	User_id int `json:"user_id"`
}

// prompt_answers_min_order_by represents the prompt_answers_min_order_by GraphQL type
type prompt_answers_min_order_by struct {
}

// prompt_answers_mutation_response represents the prompt_answers_mutation_response GraphQL type
type prompt_answers_mutation_response struct {
	Affected_rows int `json:"affected_rows"`
	Returning []*prompt_answers `json:"returning"`
}

// prompt_answers_order_by represents the prompt_answers_order_by GraphQL type
type prompt_answers_order_by struct {
}

// prompt_answers_pk_columns_input represents the prompt_answers_pk_columns_input GraphQL type
type prompt_answers_pk_columns_input struct {
}

// prompt_answers_select_column represents the prompt_answers_select_column GraphQL type
type prompt_answers_select_column struct {
}

// prompt_answers_set_input represents the prompt_answers_set_input GraphQL type
type prompt_answers_set_input struct {
}

// prompt_answers_stddev_fields represents the prompt_answers_stddev_fields GraphQL type
type prompt_answers_stddev_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Original_book_id float64 `json:"original_book_id"`
	Prompt_id float64 `json:"prompt_id"`
	User_id float64 `json:"user_id"`
}

// prompt_answers_stddev_order_by represents the prompt_answers_stddev_order_by GraphQL type
type prompt_answers_stddev_order_by struct {
}

// prompt_answers_stddev_pop_fields represents the prompt_answers_stddev_pop_fields GraphQL type
type prompt_answers_stddev_pop_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Original_book_id float64 `json:"original_book_id"`
	Prompt_id float64 `json:"prompt_id"`
	User_id float64 `json:"user_id"`
}

// prompt_answers_stddev_pop_order_by represents the prompt_answers_stddev_pop_order_by GraphQL type
type prompt_answers_stddev_pop_order_by struct {
}

// prompt_answers_stddev_samp_fields represents the prompt_answers_stddev_samp_fields GraphQL type
type prompt_answers_stddev_samp_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Original_book_id float64 `json:"original_book_id"`
	Prompt_id float64 `json:"prompt_id"`
	User_id float64 `json:"user_id"`
}

// prompt_answers_stddev_samp_order_by represents the prompt_answers_stddev_samp_order_by GraphQL type
type prompt_answers_stddev_samp_order_by struct {
}

// prompt_answers_stream_cursor_input represents the prompt_answers_stream_cursor_input GraphQL type
type prompt_answers_stream_cursor_input struct {
}

// prompt_answers_stream_cursor_value_input represents the prompt_answers_stream_cursor_value_input GraphQL type
type prompt_answers_stream_cursor_value_input struct {
}

// prompt_answers_sum_fields represents the prompt_answers_sum_fields GraphQL type
type prompt_answers_sum_fields struct {
	Book_id int `json:"book_id"`
	ID int `json:"id"`
	Original_book_id int `json:"original_book_id"`
	Prompt_id int `json:"prompt_id"`
	User_id int `json:"user_id"`
}

// prompt_answers_sum_order_by represents the prompt_answers_sum_order_by GraphQL type
type prompt_answers_sum_order_by struct {
}

// prompt_answers_updates represents the prompt_answers_updates GraphQL type
type prompt_answers_updates struct {
}

// prompt_answers_var_pop_fields represents the prompt_answers_var_pop_fields GraphQL type
type prompt_answers_var_pop_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Original_book_id float64 `json:"original_book_id"`
	Prompt_id float64 `json:"prompt_id"`
	User_id float64 `json:"user_id"`
}

// prompt_answers_var_pop_order_by represents the prompt_answers_var_pop_order_by GraphQL type
type prompt_answers_var_pop_order_by struct {
}

// prompt_answers_var_samp_fields represents the prompt_answers_var_samp_fields GraphQL type
type prompt_answers_var_samp_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Original_book_id float64 `json:"original_book_id"`
	Prompt_id float64 `json:"prompt_id"`
	User_id float64 `json:"user_id"`
}

// prompt_answers_var_samp_order_by represents the prompt_answers_var_samp_order_by GraphQL type
type prompt_answers_var_samp_order_by struct {
}

// prompt_answers_variance_fields represents the prompt_answers_variance_fields GraphQL type
type prompt_answers_variance_fields struct {
	Book_id float64 `json:"book_id"`
	ID float64 `json:"id"`
	Original_book_id float64 `json:"original_book_id"`
	Prompt_id float64 `json:"prompt_id"`
	User_id float64 `json:"user_id"`
}

// prompt_answers_variance_order_by represents the prompt_answers_variance_order_by GraphQL type
type prompt_answers_variance_order_by struct {
}

// prompt_books_summary represents the prompt_books_summary GraphQL type
type prompt_books_summary struct {
	Answers_count *bigint `json:"answers_count"`
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Prompt *prompts `json:"prompt"`
	Prompt_id int `json:"prompt_id"`
}

// prompt_books_summary_aggregate_order_by represents the prompt_books_summary_aggregate_order_by GraphQL type
type prompt_books_summary_aggregate_order_by struct {
}

// prompt_books_summary_avg_order_by represents the prompt_books_summary_avg_order_by GraphQL type
type prompt_books_summary_avg_order_by struct {
}

// prompt_books_summary_bool_exp represents the prompt_books_summary_bool_exp GraphQL type
type prompt_books_summary_bool_exp struct {
}

// prompt_books_summary_max_order_by represents the prompt_books_summary_max_order_by GraphQL type
type prompt_books_summary_max_order_by struct {
}

// prompt_books_summary_min_order_by represents the prompt_books_summary_min_order_by GraphQL type
type prompt_books_summary_min_order_by struct {
}

// prompt_books_summary_order_by represents the prompt_books_summary_order_by GraphQL type
type prompt_books_summary_order_by struct {
}

// prompt_books_summary_select_column represents the prompt_books_summary_select_column GraphQL type
type prompt_books_summary_select_column struct {
}

// prompt_books_summary_stddev_order_by represents the prompt_books_summary_stddev_order_by GraphQL type
type prompt_books_summary_stddev_order_by struct {
}

// prompt_books_summary_stddev_pop_order_by represents the prompt_books_summary_stddev_pop_order_by GraphQL type
type prompt_books_summary_stddev_pop_order_by struct {
}

// prompt_books_summary_stddev_samp_order_by represents the prompt_books_summary_stddev_samp_order_by GraphQL type
type prompt_books_summary_stddev_samp_order_by struct {
}

// prompt_books_summary_stream_cursor_input represents the prompt_books_summary_stream_cursor_input GraphQL type
type prompt_books_summary_stream_cursor_input struct {
}

// prompt_books_summary_stream_cursor_value_input represents the prompt_books_summary_stream_cursor_value_input GraphQL type
type prompt_books_summary_stream_cursor_value_input struct {
}

// prompt_books_summary_sum_order_by represents the prompt_books_summary_sum_order_by GraphQL type
type prompt_books_summary_sum_order_by struct {
}

// prompt_books_summary_var_pop_order_by represents the prompt_books_summary_var_pop_order_by GraphQL type
type prompt_books_summary_var_pop_order_by struct {
}

// prompt_books_summary_var_samp_order_by represents the prompt_books_summary_var_samp_order_by GraphQL type
type prompt_books_summary_var_samp_order_by struct {
}

// prompt_books_summary_variance_order_by represents the prompt_books_summary_variance_order_by GraphQL type
type prompt_books_summary_variance_order_by struct {
}

// prompts represents the prompts GraphQL type
type prompts struct {
	Answers_count int `json:"answers_count"`
	Books_count int `json:"books_count"`
	Created_at *timestamptz `json:"created_at"`
	Description string `json:"description"`
	Featured bool `json:"featured"`
	Followed_prompts []*followed_prompts `json:"followed_prompts"`
	Followers []*followed_users `json:"followers"`
	ID int `json:"id"`
	Privacy_setting *privacy_settings `json:"privacy_setting"`
	Privacy_setting_id int `json:"privacy_setting_id"`
	Prompt_answers []*prompt_answers `json:"prompt_answers"`
	Prompt_answers_aggregate *prompt_answers_aggregate `json:"prompt_answers_aggregate"`
	Prompt_books []*prompt_books_summary `json:"prompt_books"`
	Question string `json:"question"`
	Slug string `json:"slug"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
	Users_count int `json:"users_count"`
}

// prompts_aggregate_order_by represents the prompts_aggregate_order_by GraphQL type
type prompts_aggregate_order_by struct {
}

// prompts_avg_order_by represents the prompts_avg_order_by GraphQL type
type prompts_avg_order_by struct {
}

// prompts_bool_exp represents the prompts_bool_exp GraphQL type
type prompts_bool_exp struct {
}

// prompts_max_order_by represents the prompts_max_order_by GraphQL type
type prompts_max_order_by struct {
}

// prompts_min_order_by represents the prompts_min_order_by GraphQL type
type prompts_min_order_by struct {
}

// prompts_mutation_response represents the prompts_mutation_response GraphQL type
type prompts_mutation_response struct {
	Affected_rows int `json:"affected_rows"`
	Returning []*prompts `json:"returning"`
}

// prompts_order_by represents the prompts_order_by GraphQL type
type prompts_order_by struct {
}

// prompts_select_column represents the prompts_select_column GraphQL type
type prompts_select_column struct {
}

// prompts_stddev_order_by represents the prompts_stddev_order_by GraphQL type
type prompts_stddev_order_by struct {
}

// prompts_stddev_pop_order_by represents the prompts_stddev_pop_order_by GraphQL type
type prompts_stddev_pop_order_by struct {
}

// prompts_stddev_samp_order_by represents the prompts_stddev_samp_order_by GraphQL type
type prompts_stddev_samp_order_by struct {
}

// prompts_stream_cursor_input represents the prompts_stream_cursor_input GraphQL type
type prompts_stream_cursor_input struct {
}

// prompts_stream_cursor_value_input represents the prompts_stream_cursor_value_input GraphQL type
type prompts_stream_cursor_value_input struct {
}

// prompts_sum_order_by represents the prompts_sum_order_by GraphQL type
type prompts_sum_order_by struct {
}

// prompts_var_pop_order_by represents the prompts_var_pop_order_by GraphQL type
type prompts_var_pop_order_by struct {
}

// prompts_var_samp_order_by represents the prompts_var_samp_order_by GraphQL type
type prompts_var_samp_order_by struct {
}

// prompts_variance_order_by represents the prompts_variance_order_by GraphQL type
type prompts_variance_order_by struct {
}

// publishers represents the publishers GraphQL type
type publishers struct {
	Canonical_id int `json:"canonical_id"`
	Created_at *timestamp `json:"created_at"`
	Editions []*editions `json:"editions"`
	Editions_count int `json:"editions_count"`
	ID *bigint `json:"id"`
	Locked bool `json:"locked"`
	Name string `json:"name"`
	Parent_id int `json:"parent_id"`
	Parent_publisher *publishers `json:"parent_publisher"`
	Slug string `json:"slug"`
	State string `json:"state"`
	Updated_at *timestamp `json:"updated_at"`
	User_id int `json:"user_id"`
}

// publishers_bool_exp represents the publishers_bool_exp GraphQL type
type publishers_bool_exp struct {
}

// publishers_order_by represents the publishers_order_by GraphQL type
type publishers_order_by struct {
}

// publishers_select_column represents the publishers_select_column GraphQL type
type publishers_select_column struct {
}

// publishers_stream_cursor_input represents the publishers_stream_cursor_input GraphQL type
type publishers_stream_cursor_input struct {
}

// publishers_stream_cursor_value_input represents the publishers_stream_cursor_value_input GraphQL type
type publishers_stream_cursor_value_input struct {
}

// query_root represents the query_root GraphQL type
type query_root struct {
	Activities []*activities `json:"activities"`
	Activities_by_pk *activities `json:"activities_by_pk"`
	Activity_feed []*activities `json:"activity_feed"`
	Activity_foryou_feed []*activities `json:"activity_foryou_feed"`
	Authors []*authors `json:"authors"`
	Authors_by_pk *authors `json:"authors_by_pk"`
	Book_categories []*book_categories `json:"book_categories"`
	Book_categories_by_pk *book_categories `json:"book_categories_by_pk"`
	Book_characters []*book_characters `json:"book_characters"`
	Book_characters_by_pk *book_characters `json:"book_characters_by_pk"`
	Book_collections []*book_collections `json:"book_collections"`
	Book_collections_by_pk *book_collections `json:"book_collections_by_pk"`
	Book_mappings []*book_mappings `json:"book_mappings"`
	Book_mappings_by_pk *book_mappings `json:"book_mappings_by_pk"`
	Book_series []*book_series `json:"book_series"`
	Book_series_aggregate *book_series_aggregate `json:"book_series_aggregate"`
	Book_series_by_pk *book_series `json:"book_series_by_pk"`
	Book_statuses []*book_statuses `json:"book_statuses"`
	Book_statuses_by_pk *book_statuses `json:"book_statuses_by_pk"`
	Bookles []*bookles `json:"bookles"`
	Bookles_by_pk *bookles `json:"bookles_by_pk"`
	Books []*books `json:"books"`
	Books_aggregate *books_aggregate `json:"books_aggregate"`
	Books_by_pk *books `json:"books_by_pk"`
	Books_trending *TrendingBookType `json:"books_trending"`
	Characters []*characters `json:"characters"`
	Characters_by_pk *characters `json:"characters_by_pk"`
	Collection_import_results []*collection_import_results `json:"collection_import_results"`
	Collection_import_results_by_pk *collection_import_results `json:"collection_import_results_by_pk"`
	Collection_imports []*collection_imports `json:"collection_imports"`
	Collection_imports_by_pk *collection_imports `json:"collection_imports_by_pk"`
	Contributions []*contributions `json:"contributions"`
	Contributions_aggregate *contributions_aggregate `json:"contributions_aggregate"`
	Contributions_by_pk *contributions `json:"contributions_by_pk"`
	Countries []*countries `json:"countries"`
	Countries_by_pk *countries `json:"countries_by_pk"`
	Editions []*editions `json:"editions"`
	Editions_by_pk *editions `json:"editions_by_pk"`
	Flag_statuses []*flag_statuses `json:"flag_statuses"`
	Flag_statuses_by_pk *flag_statuses `json:"flag_statuses_by_pk"`
	Followed_lists []*followed_lists `json:"followed_lists"`
	Followed_lists_by_pk *followed_lists `json:"followed_lists_by_pk"`
	Followed_prompts []*followed_prompts `json:"followed_prompts"`
	Followed_prompts_by_pk *followed_prompts `json:"followed_prompts_by_pk"`
	Followed_user_books []*followed_user_books `json:"followed_user_books"`
	Followed_user_books_aggregate *followed_user_books_aggregate `json:"followed_user_books_aggregate"`
	Followed_users []*followed_users `json:"followed_users"`
	Followed_users_by_pk *followed_users `json:"followed_users_by_pk"`
	Following_user_books []*following_user_books `json:"following_user_books"`
	Following_user_books_aggregate *following_user_books_aggregate `json:"following_user_books_aggregate"`
	Goals []*goals `json:"goals"`
	Goals_by_pk *goals `json:"goals_by_pk"`
	Images []*images `json:"images"`
	Images_by_pk *images `json:"images_by_pk"`
	Languages []*languages `json:"languages"`
	Languages_by_pk *languages `json:"languages_by_pk"`
	Likes []*likes `json:"likes"`
	Likes_by_pk *likes `json:"likes_by_pk"`
	List_books []*list_books `json:"list_books"`
	List_books_aggregate *list_books_aggregate `json:"list_books_aggregate"`
	List_books_by_pk *list_books `json:"list_books_by_pk"`
	Lists []*lists `json:"lists"`
	Lists_aggregate *lists_aggregate `json:"lists_aggregate"`
	Lists_by_pk *lists `json:"lists_by_pk"`
	Me []*users `json:"me"`
	Newsletter *NewsletterStatusType `json:"newsletter"`
	Notification_channels []*notification_channels `json:"notification_channels"`
	Notification_channels_by_pk *notification_channels `json:"notification_channels_by_pk"`
	Notification_deliveries []*notification_deliveries `json:"notification_deliveries"`
	Notification_deliveries_aggregate *notification_deliveries_aggregate `json:"notification_deliveries_aggregate"`
	Notification_deliveries_by_pk *notification_deliveries `json:"notification_deliveries_by_pk"`
	Notification_settings []*notification_settings `json:"notification_settings"`
	Notification_settings_by_pk *notification_settings `json:"notification_settings_by_pk"`
	Notification_types []*notification_types `json:"notification_types"`
	Notification_types_by_pk *notification_types `json:"notification_types_by_pk"`
	Notifications []*notifications `json:"notifications"`
	Notifications_by_pk *notifications `json:"notifications_by_pk"`
	Platforms []*platforms `json:"platforms"`
	Platforms_by_pk *platforms `json:"platforms_by_pk"`
	Privacy_settings []*privacy_settings `json:"privacy_settings"`
	Privacy_settings_by_pk *privacy_settings `json:"privacy_settings_by_pk"`
	Prompt_answers []*prompt_answers `json:"prompt_answers"`
	Prompt_answers_aggregate *prompt_answers_aggregate `json:"prompt_answers_aggregate"`
	Prompt_answers_by_pk *prompt_answers `json:"prompt_answers_by_pk"`
	Prompt_books_summary []*prompt_books_summary `json:"prompt_books_summary"`
	Prompts []*prompts `json:"prompts"`
	Prompts_by_pk *prompts `json:"prompts_by_pk"`
	Publishers []*publishers `json:"publishers"`
	Publishers_by_pk *publishers `json:"publishers_by_pk"`
	Reading_formats []*reading_formats `json:"reading_formats"`
	Reading_formats_by_pk *reading_formats `json:"reading_formats_by_pk"`
	Reading_journals []*reading_journals `json:"reading_journals"`
	Reading_journals_by_pk *reading_journals `json:"reading_journals_by_pk"`
	Reading_journals_summary []*reading_journals_summary `json:"reading_journals_summary"`
	Recommendations []*recommendations `json:"recommendations"`
	Recommendations_by_pk *recommendations `json:"recommendations_by_pk"`
	Referrals_for_user []*ReferralType `json:"referrals_for_user"`
	Search *SearchOutput `json:"search"`
	Series []*series `json:"series"`
	Series_by_pk *series `json:"series_by_pk"`
	Subscriptions *SubscriptionsType `json:"subscriptions"`
	Tag_categories []*tag_categories `json:"tag_categories"`
	Tag_categories_by_pk *tag_categories `json:"tag_categories_by_pk"`
	Taggable_counts []*taggable_counts `json:"taggable_counts"`
	Taggable_counts_by_pk *taggable_counts `json:"taggable_counts_by_pk"`
	Taggings []*taggings `json:"taggings"`
	Taggings_aggregate *taggings_aggregate `json:"taggings_aggregate"`
	Taggings_by_pk *taggings `json:"taggings_by_pk"`
	Tags []*tags `json:"tags"`
	Tags_aggregate *tags_aggregate `json:"tags_aggregate"`
	Tags_by_pk *tags `json:"tags_by_pk"`
	User_blocks []*user_blocks `json:"user_blocks"`
	User_blocks_by_pk *user_blocks `json:"user_blocks_by_pk"`
	User_book_reads []*user_book_reads `json:"user_book_reads"`
	User_book_reads_aggregate *user_book_reads_aggregate `json:"user_book_reads_aggregate"`
	User_book_reads_by_pk *user_book_reads `json:"user_book_reads_by_pk"`
	User_book_statuses []*user_book_statuses `json:"user_book_statuses"`
	User_book_statuses_aggregate *user_book_statuses_aggregate `json:"user_book_statuses_aggregate"`
	User_book_statuses_by_pk *user_book_statuses `json:"user_book_statuses_by_pk"`
	User_books []*user_books `json:"user_books"`
	User_books_aggregate *user_books_aggregate `json:"user_books_aggregate"`
	User_books_by_pk *user_books `json:"user_books_by_pk"`
	User_flags []*user_flags `json:"user_flags"`
	User_flags_by_pk *user_flags `json:"user_flags_by_pk"`
	User_referrals []*user_referrals `json:"user_referrals"`
	User_referrals_by_pk *user_referrals `json:"user_referrals_by_pk"`
	User_statuses []*user_statuses `json:"user_statuses"`
	User_statuses_by_pk *user_statuses `json:"user_statuses_by_pk"`
	Users []*users `json:"users"`
	Users_aggregate_by_created_at_date []*users_aggregate_by_created_at_date `json:"users_aggregate_by_created_at_date"`
	Users_by_pk *users `json:"users_by_pk"`
}

// reading_formats represents the reading_formats GraphQL type
type reading_formats struct {
	Format string `json:"format"`
	ID int `json:"id"`
}

// reading_formats_bool_exp represents the reading_formats_bool_exp GraphQL type
type reading_formats_bool_exp struct {
}

// reading_formats_order_by represents the reading_formats_order_by GraphQL type
type reading_formats_order_by struct {
}

// reading_formats_select_column represents the reading_formats_select_column GraphQL type
type reading_formats_select_column struct {
}

// reading_formats_stream_cursor_input represents the reading_formats_stream_cursor_input GraphQL type
type reading_formats_stream_cursor_input struct {
}

// reading_formats_stream_cursor_value_input represents the reading_formats_stream_cursor_value_input GraphQL type
type reading_formats_stream_cursor_value_input struct {
}

// reading_journals represents the reading_journals GraphQL type
type reading_journals struct {
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Created_at *timestamp `json:"created_at"`
	Edition *editions `json:"edition"`
	Edition_id int `json:"edition_id"`
	Entry string `json:"entry"`
	Event string `json:"event"`
	Followers []*followed_users `json:"followers"`
	ID *bigint `json:"id"`
	Likes []*likes `json:"likes"`
	Likes_count int `json:"likes_count"`
	Metadata *json.RawMessage `json:"metadata"`
	Object_type string `json:"object_type"`
	Privacy_setting_id int `json:"privacy_setting_id"`
	Taggings []*taggings `json:"taggings"`
	Taggings_aggregate *taggings_aggregate `json:"taggings_aggregate"`
	Updated_at *timestamp `json:"updated_at"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// reading_journals_aggregate_order_by represents the reading_journals_aggregate_order_by GraphQL type
type reading_journals_aggregate_order_by struct {
}

// reading_journals_avg_order_by represents the reading_journals_avg_order_by GraphQL type
type reading_journals_avg_order_by struct {
}

// reading_journals_bool_exp represents the reading_journals_bool_exp GraphQL type
type reading_journals_bool_exp struct {
}

// reading_journals_max_order_by represents the reading_journals_max_order_by GraphQL type
type reading_journals_max_order_by struct {
}

// reading_journals_min_order_by represents the reading_journals_min_order_by GraphQL type
type reading_journals_min_order_by struct {
}

// reading_journals_order_by represents the reading_journals_order_by GraphQL type
type reading_journals_order_by struct {
}

// reading_journals_select_column represents the reading_journals_select_column GraphQL type
type reading_journals_select_column struct {
}

// reading_journals_stddev_order_by represents the reading_journals_stddev_order_by GraphQL type
type reading_journals_stddev_order_by struct {
}

// reading_journals_stddev_pop_order_by represents the reading_journals_stddev_pop_order_by GraphQL type
type reading_journals_stddev_pop_order_by struct {
}

// reading_journals_stddev_samp_order_by represents the reading_journals_stddev_samp_order_by GraphQL type
type reading_journals_stddev_samp_order_by struct {
}

// reading_journals_stream_cursor_input represents the reading_journals_stream_cursor_input GraphQL type
type reading_journals_stream_cursor_input struct {
}

// reading_journals_stream_cursor_value_input represents the reading_journals_stream_cursor_value_input GraphQL type
type reading_journals_stream_cursor_value_input struct {
}

// reading_journals_sum_order_by represents the reading_journals_sum_order_by GraphQL type
type reading_journals_sum_order_by struct {
}

// reading_journals_summary represents the reading_journals_summary GraphQL type
type reading_journals_summary struct {
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Followers []*followed_users `json:"followers"`
	Journals_count *bigint `json:"journals_count"`
	Last_updated_at *timestamp `json:"last_updated_at"`
	Reading_journals []*reading_journals `json:"reading_journals"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// reading_journals_summary_bool_exp represents the reading_journals_summary_bool_exp GraphQL type
type reading_journals_summary_bool_exp struct {
}

// reading_journals_summary_order_by represents the reading_journals_summary_order_by GraphQL type
type reading_journals_summary_order_by struct {
}

// reading_journals_summary_select_column represents the reading_journals_summary_select_column GraphQL type
type reading_journals_summary_select_column struct {
}

// reading_journals_summary_stream_cursor_input represents the reading_journals_summary_stream_cursor_input GraphQL type
type reading_journals_summary_stream_cursor_input struct {
}

// reading_journals_summary_stream_cursor_value_input represents the reading_journals_summary_stream_cursor_value_input GraphQL type
type reading_journals_summary_stream_cursor_value_input struct {
}

// reading_journals_var_pop_order_by represents the reading_journals_var_pop_order_by GraphQL type
type reading_journals_var_pop_order_by struct {
}

// reading_journals_var_samp_order_by represents the reading_journals_var_samp_order_by GraphQL type
type reading_journals_var_samp_order_by struct {
}

// reading_journals_variance_order_by represents the reading_journals_variance_order_by GraphQL type
type reading_journals_variance_order_by struct {
}

// recommendations represents the recommendations GraphQL type
type recommendations struct {
	Context string `json:"context"`
	Created_at *timestamp `json:"created_at"`
	ID *bigint `json:"id"`
	Item_book *books `json:"item_book"`
	Item_id *bigint `json:"item_id"`
	Item_type string `json:"item_type"`
	Item_user *users `json:"item_user"`
	Score *float8 `json:"score"`
	Subject_id *bigint `json:"subject_id"`
	Subject_type string `json:"subject_type"`
	Subject_user *users `json:"subject_user"`
	Updated_at *timestamp `json:"updated_at"`
}

// recommendations_aggregate_order_by represents the recommendations_aggregate_order_by GraphQL type
type recommendations_aggregate_order_by struct {
}

// recommendations_avg_order_by represents the recommendations_avg_order_by GraphQL type
type recommendations_avg_order_by struct {
}

// recommendations_bool_exp represents the recommendations_bool_exp GraphQL type
type recommendations_bool_exp struct {
}

// recommendations_max_order_by represents the recommendations_max_order_by GraphQL type
type recommendations_max_order_by struct {
}

// recommendations_min_order_by represents the recommendations_min_order_by GraphQL type
type recommendations_min_order_by struct {
}

// recommendations_order_by represents the recommendations_order_by GraphQL type
type recommendations_order_by struct {
}

// recommendations_select_column represents the recommendations_select_column GraphQL type
type recommendations_select_column struct {
}

// recommendations_stddev_order_by represents the recommendations_stddev_order_by GraphQL type
type recommendations_stddev_order_by struct {
}

// recommendations_stddev_pop_order_by represents the recommendations_stddev_pop_order_by GraphQL type
type recommendations_stddev_pop_order_by struct {
}

// recommendations_stddev_samp_order_by represents the recommendations_stddev_samp_order_by GraphQL type
type recommendations_stddev_samp_order_by struct {
}

// recommendations_stream_cursor_input represents the recommendations_stream_cursor_input GraphQL type
type recommendations_stream_cursor_input struct {
}

// recommendations_stream_cursor_value_input represents the recommendations_stream_cursor_value_input GraphQL type
type recommendations_stream_cursor_value_input struct {
}

// recommendations_sum_order_by represents the recommendations_sum_order_by GraphQL type
type recommendations_sum_order_by struct {
}

// recommendations_var_pop_order_by represents the recommendations_var_pop_order_by GraphQL type
type recommendations_var_pop_order_by struct {
}

// recommendations_var_samp_order_by represents the recommendations_var_samp_order_by GraphQL type
type recommendations_var_samp_order_by struct {
}

// recommendations_variance_order_by represents the recommendations_variance_order_by GraphQL type
type recommendations_variance_order_by struct {
}

// series represents the series GraphQL type
type series struct {
	Author *authors `json:"author"`
	Author_id int `json:"author_id"`
	Book_series []*book_series `json:"book_series"`
	Book_series_aggregate *book_series_aggregate `json:"book_series_aggregate"`
	Books_count int `json:"books_count"`
	Canonical *series `json:"canonical"`
	Canonical_id int `json:"canonical_id"`
	Creator *users `json:"creator"`
	Description string `json:"description"`
	ID int `json:"id"`
	Identifiers *json.RawMessage `json:"identifiers"`
	Is_completed bool `json:"is_completed"`
	Locked bool `json:"locked"`
	Name string `json:"name"`
	Primary_books_count int `json:"primary_books_count"`
	Slug string `json:"slug"`
	State string `json:"state"`
	User_id int `json:"user_id"`
}

// series_bool_exp represents the series_bool_exp GraphQL type
type series_bool_exp struct {
}

// series_order_by represents the series_order_by GraphQL type
type series_order_by struct {
}

// series_select_column represents the series_select_column GraphQL type
type series_select_column struct {
}

// series_stream_cursor_input represents the series_stream_cursor_input GraphQL type
type series_stream_cursor_input struct {
}

// series_stream_cursor_value_input represents the series_stream_cursor_value_input GraphQL type
type series_stream_cursor_value_input struct {
}

// smallint represents the smallint GraphQL type
type smallint struct {
}

// smallint_comparison_exp represents the smallint_comparison_exp GraphQL type
type smallint_comparison_exp struct {
}

// subscription_root represents the subscription_root GraphQL type
type subscription_root struct {
	Activities []*activities `json:"activities"`
	Activities_by_pk *activities `json:"activities_by_pk"`
	Activities_stream []*activities `json:"activities_stream"`
	Activity_feed []*activities `json:"activity_feed"`
	Activity_foryou_feed []*activities `json:"activity_foryou_feed"`
	Authors []*authors `json:"authors"`
	Authors_by_pk *authors `json:"authors_by_pk"`
	Authors_stream []*authors `json:"authors_stream"`
	Book_categories []*book_categories `json:"book_categories"`
	Book_categories_by_pk *book_categories `json:"book_categories_by_pk"`
	Book_categories_stream []*book_categories `json:"book_categories_stream"`
	Book_characters []*book_characters `json:"book_characters"`
	Book_characters_by_pk *book_characters `json:"book_characters_by_pk"`
	Book_characters_stream []*book_characters `json:"book_characters_stream"`
	Book_collections []*book_collections `json:"book_collections"`
	Book_collections_by_pk *book_collections `json:"book_collections_by_pk"`
	Book_collections_stream []*book_collections `json:"book_collections_stream"`
	Book_mappings []*book_mappings `json:"book_mappings"`
	Book_mappings_by_pk *book_mappings `json:"book_mappings_by_pk"`
	Book_mappings_stream []*book_mappings `json:"book_mappings_stream"`
	Book_series []*book_series `json:"book_series"`
	Book_series_aggregate *book_series_aggregate `json:"book_series_aggregate"`
	Book_series_by_pk *book_series `json:"book_series_by_pk"`
	Book_series_stream []*book_series `json:"book_series_stream"`
	Book_statuses []*book_statuses `json:"book_statuses"`
	Book_statuses_by_pk *book_statuses `json:"book_statuses_by_pk"`
	Book_statuses_stream []*book_statuses `json:"book_statuses_stream"`
	Bookles []*bookles `json:"bookles"`
	Bookles_by_pk *bookles `json:"bookles_by_pk"`
	Bookles_stream []*bookles `json:"bookles_stream"`
	Books []*books `json:"books"`
	Books_aggregate *books_aggregate `json:"books_aggregate"`
	Books_by_pk *books `json:"books_by_pk"`
	Books_stream []*books `json:"books_stream"`
	Characters []*characters `json:"characters"`
	Characters_by_pk *characters `json:"characters_by_pk"`
	Characters_stream []*characters `json:"characters_stream"`
	Collection_import_results []*collection_import_results `json:"collection_import_results"`
	Collection_import_results_by_pk *collection_import_results `json:"collection_import_results_by_pk"`
	Collection_import_results_stream []*collection_import_results `json:"collection_import_results_stream"`
	Collection_imports []*collection_imports `json:"collection_imports"`
	Collection_imports_by_pk *collection_imports `json:"collection_imports_by_pk"`
	Collection_imports_stream []*collection_imports `json:"collection_imports_stream"`
	Contributions []*contributions `json:"contributions"`
	Contributions_aggregate *contributions_aggregate `json:"contributions_aggregate"`
	Contributions_by_pk *contributions `json:"contributions_by_pk"`
	Contributions_stream []*contributions `json:"contributions_stream"`
	Countries []*countries `json:"countries"`
	Countries_by_pk *countries `json:"countries_by_pk"`
	Countries_stream []*countries `json:"countries_stream"`
	Editions []*editions `json:"editions"`
	Editions_by_pk *editions `json:"editions_by_pk"`
	Editions_stream []*editions `json:"editions_stream"`
	Flag_statuses []*flag_statuses `json:"flag_statuses"`
	Flag_statuses_by_pk *flag_statuses `json:"flag_statuses_by_pk"`
	Flag_statuses_stream []*flag_statuses `json:"flag_statuses_stream"`
	Followed_lists []*followed_lists `json:"followed_lists"`
	Followed_lists_by_pk *followed_lists `json:"followed_lists_by_pk"`
	Followed_lists_stream []*followed_lists `json:"followed_lists_stream"`
	Followed_prompts []*followed_prompts `json:"followed_prompts"`
	Followed_prompts_by_pk *followed_prompts `json:"followed_prompts_by_pk"`
	Followed_prompts_stream []*followed_prompts `json:"followed_prompts_stream"`
	Followed_user_books []*followed_user_books `json:"followed_user_books"`
	Followed_user_books_aggregate *followed_user_books_aggregate `json:"followed_user_books_aggregate"`
	Followed_user_books_stream []*followed_user_books `json:"followed_user_books_stream"`
	Followed_users []*followed_users `json:"followed_users"`
	Followed_users_by_pk *followed_users `json:"followed_users_by_pk"`
	Followed_users_stream []*followed_users `json:"followed_users_stream"`
	Following_user_books []*following_user_books `json:"following_user_books"`
	Following_user_books_aggregate *following_user_books_aggregate `json:"following_user_books_aggregate"`
	Following_user_books_stream []*following_user_books `json:"following_user_books_stream"`
	Goals []*goals `json:"goals"`
	Goals_by_pk *goals `json:"goals_by_pk"`
	Goals_stream []*goals `json:"goals_stream"`
	Images []*images `json:"images"`
	Images_by_pk *images `json:"images_by_pk"`
	Images_stream []*images `json:"images_stream"`
	Languages []*languages `json:"languages"`
	Languages_by_pk *languages `json:"languages_by_pk"`
	Languages_stream []*languages `json:"languages_stream"`
	Likes []*likes `json:"likes"`
	Likes_by_pk *likes `json:"likes_by_pk"`
	Likes_stream []*likes `json:"likes_stream"`
	List_books []*list_books `json:"list_books"`
	List_books_aggregate *list_books_aggregate `json:"list_books_aggregate"`
	List_books_by_pk *list_books `json:"list_books_by_pk"`
	List_books_stream []*list_books `json:"list_books_stream"`
	Lists []*lists `json:"lists"`
	Lists_aggregate *lists_aggregate `json:"lists_aggregate"`
	Lists_by_pk *lists `json:"lists_by_pk"`
	Lists_stream []*lists `json:"lists_stream"`
	Me []*users `json:"me"`
	Notification_channels []*notification_channels `json:"notification_channels"`
	Notification_channels_by_pk *notification_channels `json:"notification_channels_by_pk"`
	Notification_channels_stream []*notification_channels `json:"notification_channels_stream"`
	Notification_deliveries []*notification_deliveries `json:"notification_deliveries"`
	Notification_deliveries_aggregate *notification_deliveries_aggregate `json:"notification_deliveries_aggregate"`
	Notification_deliveries_by_pk *notification_deliveries `json:"notification_deliveries_by_pk"`
	Notification_deliveries_stream []*notification_deliveries `json:"notification_deliveries_stream"`
	Notification_settings []*notification_settings `json:"notification_settings"`
	Notification_settings_by_pk *notification_settings `json:"notification_settings_by_pk"`
	Notification_settings_stream []*notification_settings `json:"notification_settings_stream"`
	Notification_types []*notification_types `json:"notification_types"`
	Notification_types_by_pk *notification_types `json:"notification_types_by_pk"`
	Notification_types_stream []*notification_types `json:"notification_types_stream"`
	Notifications []*notifications `json:"notifications"`
	Notifications_by_pk *notifications `json:"notifications_by_pk"`
	Notifications_stream []*notifications `json:"notifications_stream"`
	Platforms []*platforms `json:"platforms"`
	Platforms_by_pk *platforms `json:"platforms_by_pk"`
	Platforms_stream []*platforms `json:"platforms_stream"`
	Privacy_settings []*privacy_settings `json:"privacy_settings"`
	Privacy_settings_by_pk *privacy_settings `json:"privacy_settings_by_pk"`
	Privacy_settings_stream []*privacy_settings `json:"privacy_settings_stream"`
	Prompt_answers []*prompt_answers `json:"prompt_answers"`
	Prompt_answers_aggregate *prompt_answers_aggregate `json:"prompt_answers_aggregate"`
	Prompt_answers_by_pk *prompt_answers `json:"prompt_answers_by_pk"`
	Prompt_answers_stream []*prompt_answers `json:"prompt_answers_stream"`
	Prompt_books_summary []*prompt_books_summary `json:"prompt_books_summary"`
	Prompt_books_summary_stream []*prompt_books_summary `json:"prompt_books_summary_stream"`
	Prompts []*prompts `json:"prompts"`
	Prompts_by_pk *prompts `json:"prompts_by_pk"`
	Prompts_stream []*prompts `json:"prompts_stream"`
	Publishers []*publishers `json:"publishers"`
	Publishers_by_pk *publishers `json:"publishers_by_pk"`
	Publishers_stream []*publishers `json:"publishers_stream"`
	Reading_formats []*reading_formats `json:"reading_formats"`
	Reading_formats_by_pk *reading_formats `json:"reading_formats_by_pk"`
	Reading_formats_stream []*reading_formats `json:"reading_formats_stream"`
	Reading_journals []*reading_journals `json:"reading_journals"`
	Reading_journals_by_pk *reading_journals `json:"reading_journals_by_pk"`
	Reading_journals_stream []*reading_journals `json:"reading_journals_stream"`
	Reading_journals_summary []*reading_journals_summary `json:"reading_journals_summary"`
	Reading_journals_summary_stream []*reading_journals_summary `json:"reading_journals_summary_stream"`
	Recommendations []*recommendations `json:"recommendations"`
	Recommendations_by_pk *recommendations `json:"recommendations_by_pk"`
	Recommendations_stream []*recommendations `json:"recommendations_stream"`
	Series []*series `json:"series"`
	Series_by_pk *series `json:"series_by_pk"`
	Series_stream []*series `json:"series_stream"`
	Tag_categories []*tag_categories `json:"tag_categories"`
	Tag_categories_by_pk *tag_categories `json:"tag_categories_by_pk"`
	Tag_categories_stream []*tag_categories `json:"tag_categories_stream"`
	Taggable_counts []*taggable_counts `json:"taggable_counts"`
	Taggable_counts_by_pk *taggable_counts `json:"taggable_counts_by_pk"`
	Taggable_counts_stream []*taggable_counts `json:"taggable_counts_stream"`
	Taggings []*taggings `json:"taggings"`
	Taggings_aggregate *taggings_aggregate `json:"taggings_aggregate"`
	Taggings_by_pk *taggings `json:"taggings_by_pk"`
	Taggings_stream []*taggings `json:"taggings_stream"`
	Tags []*tags `json:"tags"`
	Tags_aggregate *tags_aggregate `json:"tags_aggregate"`
	Tags_by_pk *tags `json:"tags_by_pk"`
	Tags_stream []*tags `json:"tags_stream"`
	User_blocks []*user_blocks `json:"user_blocks"`
	User_blocks_by_pk *user_blocks `json:"user_blocks_by_pk"`
	User_blocks_stream []*user_blocks `json:"user_blocks_stream"`
	User_book_reads []*user_book_reads `json:"user_book_reads"`
	User_book_reads_aggregate *user_book_reads_aggregate `json:"user_book_reads_aggregate"`
	User_book_reads_by_pk *user_book_reads `json:"user_book_reads_by_pk"`
	User_book_reads_stream []*user_book_reads `json:"user_book_reads_stream"`
	User_book_statuses []*user_book_statuses `json:"user_book_statuses"`
	User_book_statuses_aggregate *user_book_statuses_aggregate `json:"user_book_statuses_aggregate"`
	User_book_statuses_by_pk *user_book_statuses `json:"user_book_statuses_by_pk"`
	User_book_statuses_stream []*user_book_statuses `json:"user_book_statuses_stream"`
	User_books []*user_books `json:"user_books"`
	User_books_aggregate *user_books_aggregate `json:"user_books_aggregate"`
	User_books_by_pk *user_books `json:"user_books_by_pk"`
	User_books_stream []*user_books `json:"user_books_stream"`
	User_flags []*user_flags `json:"user_flags"`
	User_flags_by_pk *user_flags `json:"user_flags_by_pk"`
	User_flags_stream []*user_flags `json:"user_flags_stream"`
	User_referrals []*user_referrals `json:"user_referrals"`
	User_referrals_by_pk *user_referrals `json:"user_referrals_by_pk"`
	User_referrals_stream []*user_referrals `json:"user_referrals_stream"`
	User_statuses []*user_statuses `json:"user_statuses"`
	User_statuses_by_pk *user_statuses `json:"user_statuses_by_pk"`
	User_statuses_stream []*user_statuses `json:"user_statuses_stream"`
	Users []*users `json:"users"`
	Users_aggregate_by_created_at_date []*users_aggregate_by_created_at_date `json:"users_aggregate_by_created_at_date"`
	Users_aggregate_by_created_at_date_stream []*users_aggregate_by_created_at_date `json:"users_aggregate_by_created_at_date_stream"`
	Users_by_pk *users `json:"users_by_pk"`
	Users_stream []*users `json:"users_stream"`
}

// tag_categories represents the tag_categories GraphQL type
type tag_categories struct {
	Category string `json:"category"`
	Created_at *timestamp `json:"created_at"`
	ID *bigint `json:"id"`
	Slug string `json:"slug"`
	Tags []*tags `json:"tags"`
	Tags_aggregate *tags_aggregate `json:"tags_aggregate"`
}

// tag_categories_bool_exp represents the tag_categories_bool_exp GraphQL type
type tag_categories_bool_exp struct {
}

// tag_categories_order_by represents the tag_categories_order_by GraphQL type
type tag_categories_order_by struct {
}

// tag_categories_select_column represents the tag_categories_select_column GraphQL type
type tag_categories_select_column struct {
}

// tag_categories_stream_cursor_input represents the tag_categories_stream_cursor_input GraphQL type
type tag_categories_stream_cursor_input struct {
}

// tag_categories_stream_cursor_value_input represents the tag_categories_stream_cursor_value_input GraphQL type
type tag_categories_stream_cursor_value_input struct {
}

// taggable_counts represents the taggable_counts GraphQL type
type taggable_counts struct {
	Book *books `json:"book"`
	Count int `json:"count"`
	Created_at *timestamp `json:"created_at"`
	Hardcover_tagged bool `json:"hardcover_tagged"`
	ID *bigint `json:"id"`
	Spoiler_ratio *float8 `json:"spoiler_ratio"`
	Tag *tags `json:"tag"`
	Tag_id int `json:"tag_id"`
	Taggable_id *bigint `json:"taggable_id"`
	Taggable_type string `json:"taggable_type"`
	Updated_at *timestamp `json:"updated_at"`
}

// taggable_counts_aggregate_order_by represents the taggable_counts_aggregate_order_by GraphQL type
type taggable_counts_aggregate_order_by struct {
}

// taggable_counts_avg_order_by represents the taggable_counts_avg_order_by GraphQL type
type taggable_counts_avg_order_by struct {
}

// taggable_counts_bool_exp represents the taggable_counts_bool_exp GraphQL type
type taggable_counts_bool_exp struct {
}

// taggable_counts_max_order_by represents the taggable_counts_max_order_by GraphQL type
type taggable_counts_max_order_by struct {
}

// taggable_counts_min_order_by represents the taggable_counts_min_order_by GraphQL type
type taggable_counts_min_order_by struct {
}

// taggable_counts_order_by represents the taggable_counts_order_by GraphQL type
type taggable_counts_order_by struct {
}

// taggable_counts_select_column represents the taggable_counts_select_column GraphQL type
type taggable_counts_select_column struct {
}

// taggable_counts_stddev_order_by represents the taggable_counts_stddev_order_by GraphQL type
type taggable_counts_stddev_order_by struct {
}

// taggable_counts_stddev_pop_order_by represents the taggable_counts_stddev_pop_order_by GraphQL type
type taggable_counts_stddev_pop_order_by struct {
}

// taggable_counts_stddev_samp_order_by represents the taggable_counts_stddev_samp_order_by GraphQL type
type taggable_counts_stddev_samp_order_by struct {
}

// taggable_counts_stream_cursor_input represents the taggable_counts_stream_cursor_input GraphQL type
type taggable_counts_stream_cursor_input struct {
}

// taggable_counts_stream_cursor_value_input represents the taggable_counts_stream_cursor_value_input GraphQL type
type taggable_counts_stream_cursor_value_input struct {
}

// taggable_counts_sum_order_by represents the taggable_counts_sum_order_by GraphQL type
type taggable_counts_sum_order_by struct {
}

// taggable_counts_var_pop_order_by represents the taggable_counts_var_pop_order_by GraphQL type
type taggable_counts_var_pop_order_by struct {
}

// taggable_counts_var_samp_order_by represents the taggable_counts_var_samp_order_by GraphQL type
type taggable_counts_var_samp_order_by struct {
}

// taggable_counts_variance_order_by represents the taggable_counts_variance_order_by GraphQL type
type taggable_counts_variance_order_by struct {
}

// taggings represents the taggings GraphQL type
type taggings struct {
	Book *books `json:"book"`
	Created_at *timestamp `json:"created_at"`
	ID *bigint `json:"id"`
	Spoiler bool `json:"spoiler"`
	Tag *tags `json:"tag"`
	Tag_id int `json:"tag_id"`
	Taggable_id *bigint `json:"taggable_id"`
	Taggable_type string `json:"taggable_type"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// taggings_aggregate represents the taggings_aggregate GraphQL type
type taggings_aggregate struct {
	Aggregate *taggings_aggregate_fields `json:"aggregate"`
	Nodes []*taggings `json:"nodes"`
}

// taggings_aggregate_bool_exp represents the taggings_aggregate_bool_exp GraphQL type
type taggings_aggregate_bool_exp struct {
}

// taggings_aggregate_bool_exp_bool_and represents the taggings_aggregate_bool_exp_bool_and GraphQL type
type taggings_aggregate_bool_exp_bool_and struct {
}

// taggings_aggregate_bool_exp_bool_or represents the taggings_aggregate_bool_exp_bool_or GraphQL type
type taggings_aggregate_bool_exp_bool_or struct {
}

// taggings_aggregate_bool_exp_count represents the taggings_aggregate_bool_exp_count GraphQL type
type taggings_aggregate_bool_exp_count struct {
}

// taggings_aggregate_fields represents the taggings_aggregate_fields GraphQL type
type taggings_aggregate_fields struct {
	Avg *taggings_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *taggings_max_fields `json:"max"`
	Min *taggings_min_fields `json:"min"`
	Stddev *taggings_stddev_fields `json:"stddev"`
	Stddev_pop *taggings_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *taggings_stddev_samp_fields `json:"stddev_samp"`
	Sum *taggings_sum_fields `json:"sum"`
	Var_pop *taggings_var_pop_fields `json:"var_pop"`
	Var_samp *taggings_var_samp_fields `json:"var_samp"`
	Variance *taggings_variance_fields `json:"variance"`
}

// taggings_aggregate_order_by represents the taggings_aggregate_order_by GraphQL type
type taggings_aggregate_order_by struct {
}

// taggings_avg_fields represents the taggings_avg_fields GraphQL type
type taggings_avg_fields struct {
	ID float64 `json:"id"`
	Tag_id float64 `json:"tag_id"`
	Taggable_id float64 `json:"taggable_id"`
	User_id float64 `json:"user_id"`
}

// taggings_avg_order_by represents the taggings_avg_order_by GraphQL type
type taggings_avg_order_by struct {
}

// taggings_bool_exp represents the taggings_bool_exp GraphQL type
type taggings_bool_exp struct {
}

// taggings_max_fields represents the taggings_max_fields GraphQL type
type taggings_max_fields struct {
	Created_at *timestamp `json:"created_at"`
	ID *bigint `json:"id"`
	Tag_id int `json:"tag_id"`
	Taggable_id *bigint `json:"taggable_id"`
	Taggable_type string `json:"taggable_type"`
	User_id int `json:"user_id"`
}

// taggings_max_order_by represents the taggings_max_order_by GraphQL type
type taggings_max_order_by struct {
}

// taggings_min_fields represents the taggings_min_fields GraphQL type
type taggings_min_fields struct {
	Created_at *timestamp `json:"created_at"`
	ID *bigint `json:"id"`
	Tag_id int `json:"tag_id"`
	Taggable_id *bigint `json:"taggable_id"`
	Taggable_type string `json:"taggable_type"`
	User_id int `json:"user_id"`
}

// taggings_min_order_by represents the taggings_min_order_by GraphQL type
type taggings_min_order_by struct {
}

// taggings_order_by represents the taggings_order_by GraphQL type
type taggings_order_by struct {
}

// taggings_select_column represents the taggings_select_column GraphQL type
type taggings_select_column struct {
}

// taggings_select_column_taggings_aggregate_bool_exp_bool_and_arguments_columns represents the taggings_select_column_taggings_aggregate_bool_exp_bool_and_arguments_columns GraphQL type
type taggings_select_column_taggings_aggregate_bool_exp_bool_and_arguments_columns struct {
}

// taggings_select_column_taggings_aggregate_bool_exp_bool_or_arguments_columns represents the taggings_select_column_taggings_aggregate_bool_exp_bool_or_arguments_columns GraphQL type
type taggings_select_column_taggings_aggregate_bool_exp_bool_or_arguments_columns struct {
}

// taggings_stddev_fields represents the taggings_stddev_fields GraphQL type
type taggings_stddev_fields struct {
	ID float64 `json:"id"`
	Tag_id float64 `json:"tag_id"`
	Taggable_id float64 `json:"taggable_id"`
	User_id float64 `json:"user_id"`
}

// taggings_stddev_order_by represents the taggings_stddev_order_by GraphQL type
type taggings_stddev_order_by struct {
}

// taggings_stddev_pop_fields represents the taggings_stddev_pop_fields GraphQL type
type taggings_stddev_pop_fields struct {
	ID float64 `json:"id"`
	Tag_id float64 `json:"tag_id"`
	Taggable_id float64 `json:"taggable_id"`
	User_id float64 `json:"user_id"`
}

// taggings_stddev_pop_order_by represents the taggings_stddev_pop_order_by GraphQL type
type taggings_stddev_pop_order_by struct {
}

// taggings_stddev_samp_fields represents the taggings_stddev_samp_fields GraphQL type
type taggings_stddev_samp_fields struct {
	ID float64 `json:"id"`
	Tag_id float64 `json:"tag_id"`
	Taggable_id float64 `json:"taggable_id"`
	User_id float64 `json:"user_id"`
}

// taggings_stddev_samp_order_by represents the taggings_stddev_samp_order_by GraphQL type
type taggings_stddev_samp_order_by struct {
}

// taggings_stream_cursor_input represents the taggings_stream_cursor_input GraphQL type
type taggings_stream_cursor_input struct {
}

// taggings_stream_cursor_value_input represents the taggings_stream_cursor_value_input GraphQL type
type taggings_stream_cursor_value_input struct {
}

// taggings_sum_fields represents the taggings_sum_fields GraphQL type
type taggings_sum_fields struct {
	ID *bigint `json:"id"`
	Tag_id int `json:"tag_id"`
	Taggable_id *bigint `json:"taggable_id"`
	User_id int `json:"user_id"`
}

// taggings_sum_order_by represents the taggings_sum_order_by GraphQL type
type taggings_sum_order_by struct {
}

// taggings_var_pop_fields represents the taggings_var_pop_fields GraphQL type
type taggings_var_pop_fields struct {
	ID float64 `json:"id"`
	Tag_id float64 `json:"tag_id"`
	Taggable_id float64 `json:"taggable_id"`
	User_id float64 `json:"user_id"`
}

// taggings_var_pop_order_by represents the taggings_var_pop_order_by GraphQL type
type taggings_var_pop_order_by struct {
}

// taggings_var_samp_fields represents the taggings_var_samp_fields GraphQL type
type taggings_var_samp_fields struct {
	ID float64 `json:"id"`
	Tag_id float64 `json:"tag_id"`
	Taggable_id float64 `json:"taggable_id"`
	User_id float64 `json:"user_id"`
}

// taggings_var_samp_order_by represents the taggings_var_samp_order_by GraphQL type
type taggings_var_samp_order_by struct {
}

// taggings_variance_fields represents the taggings_variance_fields GraphQL type
type taggings_variance_fields struct {
	ID float64 `json:"id"`
	Tag_id float64 `json:"tag_id"`
	Taggable_id float64 `json:"taggable_id"`
	User_id float64 `json:"user_id"`
}

// taggings_variance_order_by represents the taggings_variance_order_by GraphQL type
type taggings_variance_order_by struct {
}

// tags represents the tags GraphQL type
type tags struct {
	Count int `json:"count"`
	ID *bigint `json:"id"`
	Slug string `json:"slug"`
	Tag string `json:"tag"`
	Tag_category *tag_categories `json:"tag_category"`
	Tag_category_id int `json:"tag_category_id"`
	Taggings []*taggings `json:"taggings"`
	Taggings_aggregate *taggings_aggregate `json:"taggings_aggregate"`
}

// tags_aggregate represents the tags_aggregate GraphQL type
type tags_aggregate struct {
	Aggregate *tags_aggregate_fields `json:"aggregate"`
	Nodes []*tags `json:"nodes"`
}

// tags_aggregate_bool_exp represents the tags_aggregate_bool_exp GraphQL type
type tags_aggregate_bool_exp struct {
}

// tags_aggregate_bool_exp_count represents the tags_aggregate_bool_exp_count GraphQL type
type tags_aggregate_bool_exp_count struct {
}

// tags_aggregate_fields represents the tags_aggregate_fields GraphQL type
type tags_aggregate_fields struct {
	Avg *tags_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *tags_max_fields `json:"max"`
	Min *tags_min_fields `json:"min"`
	Stddev *tags_stddev_fields `json:"stddev"`
	Stddev_pop *tags_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *tags_stddev_samp_fields `json:"stddev_samp"`
	Sum *tags_sum_fields `json:"sum"`
	Var_pop *tags_var_pop_fields `json:"var_pop"`
	Var_samp *tags_var_samp_fields `json:"var_samp"`
	Variance *tags_variance_fields `json:"variance"`
}

// tags_aggregate_order_by represents the tags_aggregate_order_by GraphQL type
type tags_aggregate_order_by struct {
}

// tags_avg_fields represents the tags_avg_fields GraphQL type
type tags_avg_fields struct {
	Count float64 `json:"count"`
	ID float64 `json:"id"`
	Tag_category_id float64 `json:"tag_category_id"`
}

// tags_avg_order_by represents the tags_avg_order_by GraphQL type
type tags_avg_order_by struct {
}

// tags_bool_exp represents the tags_bool_exp GraphQL type
type tags_bool_exp struct {
}

// tags_max_fields represents the tags_max_fields GraphQL type
type tags_max_fields struct {
	Count int `json:"count"`
	ID *bigint `json:"id"`
	Slug string `json:"slug"`
	Tag string `json:"tag"`
	Tag_category_id int `json:"tag_category_id"`
}

// tags_max_order_by represents the tags_max_order_by GraphQL type
type tags_max_order_by struct {
}

// tags_min_fields represents the tags_min_fields GraphQL type
type tags_min_fields struct {
	Count int `json:"count"`
	ID *bigint `json:"id"`
	Slug string `json:"slug"`
	Tag string `json:"tag"`
	Tag_category_id int `json:"tag_category_id"`
}

// tags_min_order_by represents the tags_min_order_by GraphQL type
type tags_min_order_by struct {
}

// tags_order_by represents the tags_order_by GraphQL type
type tags_order_by struct {
}

// tags_select_column represents the tags_select_column GraphQL type
type tags_select_column struct {
}

// tags_stddev_fields represents the tags_stddev_fields GraphQL type
type tags_stddev_fields struct {
	Count float64 `json:"count"`
	ID float64 `json:"id"`
	Tag_category_id float64 `json:"tag_category_id"`
}

// tags_stddev_order_by represents the tags_stddev_order_by GraphQL type
type tags_stddev_order_by struct {
}

// tags_stddev_pop_fields represents the tags_stddev_pop_fields GraphQL type
type tags_stddev_pop_fields struct {
	Count float64 `json:"count"`
	ID float64 `json:"id"`
	Tag_category_id float64 `json:"tag_category_id"`
}

// tags_stddev_pop_order_by represents the tags_stddev_pop_order_by GraphQL type
type tags_stddev_pop_order_by struct {
}

// tags_stddev_samp_fields represents the tags_stddev_samp_fields GraphQL type
type tags_stddev_samp_fields struct {
	Count float64 `json:"count"`
	ID float64 `json:"id"`
	Tag_category_id float64 `json:"tag_category_id"`
}

// tags_stddev_samp_order_by represents the tags_stddev_samp_order_by GraphQL type
type tags_stddev_samp_order_by struct {
}

// tags_stream_cursor_input represents the tags_stream_cursor_input GraphQL type
type tags_stream_cursor_input struct {
}

// tags_stream_cursor_value_input represents the tags_stream_cursor_value_input GraphQL type
type tags_stream_cursor_value_input struct {
}

// tags_sum_fields represents the tags_sum_fields GraphQL type
type tags_sum_fields struct {
	Count int `json:"count"`
	ID *bigint `json:"id"`
	Tag_category_id int `json:"tag_category_id"`
}

// tags_sum_order_by represents the tags_sum_order_by GraphQL type
type tags_sum_order_by struct {
}

// tags_var_pop_fields represents the tags_var_pop_fields GraphQL type
type tags_var_pop_fields struct {
	Count float64 `json:"count"`
	ID float64 `json:"id"`
	Tag_category_id float64 `json:"tag_category_id"`
}

// tags_var_pop_order_by represents the tags_var_pop_order_by GraphQL type
type tags_var_pop_order_by struct {
}

// tags_var_samp_fields represents the tags_var_samp_fields GraphQL type
type tags_var_samp_fields struct {
	Count float64 `json:"count"`
	ID float64 `json:"id"`
	Tag_category_id float64 `json:"tag_category_id"`
}

// tags_var_samp_order_by represents the tags_var_samp_order_by GraphQL type
type tags_var_samp_order_by struct {
}

// tags_variance_fields represents the tags_variance_fields GraphQL type
type tags_variance_fields struct {
	Count float64 `json:"count"`
	ID float64 `json:"id"`
	Tag_category_id float64 `json:"tag_category_id"`
}

// tags_variance_order_by represents the tags_variance_order_by GraphQL type
type tags_variance_order_by struct {
}

// timestamp represents the timestamp GraphQL type
type timestamp struct {
}

// timestamp_comparison_exp represents the timestamp_comparison_exp GraphQL type
type timestamp_comparison_exp struct {
}

// timestamptz represents the timestamptz GraphQL type
type timestamptz struct {
}

// timestamptz_comparison_exp represents the timestamptz_comparison_exp GraphQL type
type timestamptz_comparison_exp struct {
}

// update_user_input represents the update_user_input GraphQL type
type update_user_input struct {
}

// user_blocks represents the user_blocks GraphQL type
type user_blocks struct {
	Blocked_user *users `json:"blocked_user"`
	Blocked_user_id int `json:"blocked_user_id"`
	Created_at *timestamp `json:"created_at"`
	ID *bigint `json:"id"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// user_blocks_aggregate_order_by represents the user_blocks_aggregate_order_by GraphQL type
type user_blocks_aggregate_order_by struct {
}

// user_blocks_avg_order_by represents the user_blocks_avg_order_by GraphQL type
type user_blocks_avg_order_by struct {
}

// user_blocks_bool_exp represents the user_blocks_bool_exp GraphQL type
type user_blocks_bool_exp struct {
}

// user_blocks_constraint represents the user_blocks_constraint GraphQL type
type user_blocks_constraint struct {
}

// user_blocks_insert_input represents the user_blocks_insert_input GraphQL type
type user_blocks_insert_input struct {
}

// user_blocks_max_order_by represents the user_blocks_max_order_by GraphQL type
type user_blocks_max_order_by struct {
}

// user_blocks_min_order_by represents the user_blocks_min_order_by GraphQL type
type user_blocks_min_order_by struct {
}

// user_blocks_mutation_response represents the user_blocks_mutation_response GraphQL type
type user_blocks_mutation_response struct {
	Affected_rows int `json:"affected_rows"`
	Returning []*user_blocks `json:"returning"`
}

// user_blocks_on_conflict represents the user_blocks_on_conflict GraphQL type
type user_blocks_on_conflict struct {
}

// user_blocks_order_by represents the user_blocks_order_by GraphQL type
type user_blocks_order_by struct {
}

// user_blocks_select_column represents the user_blocks_select_column GraphQL type
type user_blocks_select_column struct {
}

// user_blocks_stddev_order_by represents the user_blocks_stddev_order_by GraphQL type
type user_blocks_stddev_order_by struct {
}

// user_blocks_stddev_pop_order_by represents the user_blocks_stddev_pop_order_by GraphQL type
type user_blocks_stddev_pop_order_by struct {
}

// user_blocks_stddev_samp_order_by represents the user_blocks_stddev_samp_order_by GraphQL type
type user_blocks_stddev_samp_order_by struct {
}

// user_blocks_stream_cursor_input represents the user_blocks_stream_cursor_input GraphQL type
type user_blocks_stream_cursor_input struct {
}

// user_blocks_stream_cursor_value_input represents the user_blocks_stream_cursor_value_input GraphQL type
type user_blocks_stream_cursor_value_input struct {
}

// user_blocks_sum_order_by represents the user_blocks_sum_order_by GraphQL type
type user_blocks_sum_order_by struct {
}

// user_blocks_update_column represents the user_blocks_update_column GraphQL type
type user_blocks_update_column struct {
}

// user_blocks_var_pop_order_by represents the user_blocks_var_pop_order_by GraphQL type
type user_blocks_var_pop_order_by struct {
}

// user_blocks_var_samp_order_by represents the user_blocks_var_samp_order_by GraphQL type
type user_blocks_var_samp_order_by struct {
}

// user_blocks_variance_order_by represents the user_blocks_variance_order_by GraphQL type
type user_blocks_variance_order_by struct {
}

// user_book_reads represents the user_book_reads GraphQL type
type user_book_reads struct {
	Edition *editions `json:"edition"`
	Edition_id int `json:"edition_id"`
	Finished_at *date `json:"finished_at"`
	ID int `json:"id"`
	Paused_at *date `json:"paused_at"`
	Progress *float8 `json:"progress"`
	Progress_pages int `json:"progress_pages"`
	Progress_seconds int `json:"progress_seconds"`
	Started_at *date `json:"started_at"`
	User_book *user_books `json:"user_book"`
	User_book_id int `json:"user_book_id"`
}

// user_book_reads_aggregate represents the user_book_reads_aggregate GraphQL type
type user_book_reads_aggregate struct {
	Aggregate *user_book_reads_aggregate_fields `json:"aggregate"`
	Nodes []*user_book_reads `json:"nodes"`
}

// user_book_reads_aggregate_bool_exp represents the user_book_reads_aggregate_bool_exp GraphQL type
type user_book_reads_aggregate_bool_exp struct {
}

// user_book_reads_aggregate_bool_exp_avg represents the user_book_reads_aggregate_bool_exp_avg GraphQL type
type user_book_reads_aggregate_bool_exp_avg struct {
}

// user_book_reads_aggregate_bool_exp_corr represents the user_book_reads_aggregate_bool_exp_corr GraphQL type
type user_book_reads_aggregate_bool_exp_corr struct {
}

// user_book_reads_aggregate_bool_exp_corr_arguments represents the user_book_reads_aggregate_bool_exp_corr_arguments GraphQL type
type user_book_reads_aggregate_bool_exp_corr_arguments struct {
}

// user_book_reads_aggregate_bool_exp_count represents the user_book_reads_aggregate_bool_exp_count GraphQL type
type user_book_reads_aggregate_bool_exp_count struct {
}

// user_book_reads_aggregate_bool_exp_covar_samp represents the user_book_reads_aggregate_bool_exp_covar_samp GraphQL type
type user_book_reads_aggregate_bool_exp_covar_samp struct {
}

// user_book_reads_aggregate_bool_exp_covar_samp_arguments represents the user_book_reads_aggregate_bool_exp_covar_samp_arguments GraphQL type
type user_book_reads_aggregate_bool_exp_covar_samp_arguments struct {
}

// user_book_reads_aggregate_bool_exp_max represents the user_book_reads_aggregate_bool_exp_max GraphQL type
type user_book_reads_aggregate_bool_exp_max struct {
}

// user_book_reads_aggregate_bool_exp_min represents the user_book_reads_aggregate_bool_exp_min GraphQL type
type user_book_reads_aggregate_bool_exp_min struct {
}

// user_book_reads_aggregate_bool_exp_stddev_samp represents the user_book_reads_aggregate_bool_exp_stddev_samp GraphQL type
type user_book_reads_aggregate_bool_exp_stddev_samp struct {
}

// user_book_reads_aggregate_bool_exp_sum represents the user_book_reads_aggregate_bool_exp_sum GraphQL type
type user_book_reads_aggregate_bool_exp_sum struct {
}

// user_book_reads_aggregate_bool_exp_var_samp represents the user_book_reads_aggregate_bool_exp_var_samp GraphQL type
type user_book_reads_aggregate_bool_exp_var_samp struct {
}

// user_book_reads_aggregate_fields represents the user_book_reads_aggregate_fields GraphQL type
type user_book_reads_aggregate_fields struct {
	Avg *user_book_reads_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *user_book_reads_max_fields `json:"max"`
	Min *user_book_reads_min_fields `json:"min"`
	Stddev *user_book_reads_stddev_fields `json:"stddev"`
	Stddev_pop *user_book_reads_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *user_book_reads_stddev_samp_fields `json:"stddev_samp"`
	Sum *user_book_reads_sum_fields `json:"sum"`
	Var_pop *user_book_reads_var_pop_fields `json:"var_pop"`
	Var_samp *user_book_reads_var_samp_fields `json:"var_samp"`
	Variance *user_book_reads_variance_fields `json:"variance"`
}

// user_book_reads_aggregate_order_by represents the user_book_reads_aggregate_order_by GraphQL type
type user_book_reads_aggregate_order_by struct {
}

// user_book_reads_avg_fields represents the user_book_reads_avg_fields GraphQL type
type user_book_reads_avg_fields struct {
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Progress float64 `json:"progress"`
	Progress_pages float64 `json:"progress_pages"`
	Progress_seconds float64 `json:"progress_seconds"`
	User_book_id float64 `json:"user_book_id"`
}

// user_book_reads_avg_order_by represents the user_book_reads_avg_order_by GraphQL type
type user_book_reads_avg_order_by struct {
}

// user_book_reads_bool_exp represents the user_book_reads_bool_exp GraphQL type
type user_book_reads_bool_exp struct {
}

// user_book_reads_max_fields represents the user_book_reads_max_fields GraphQL type
type user_book_reads_max_fields struct {
	Edition_id int `json:"edition_id"`
	Finished_at *date `json:"finished_at"`
	ID int `json:"id"`
	Paused_at *date `json:"paused_at"`
	Progress *float8 `json:"progress"`
	Progress_pages int `json:"progress_pages"`
	Progress_seconds int `json:"progress_seconds"`
	Started_at *date `json:"started_at"`
	User_book_id int `json:"user_book_id"`
}

// user_book_reads_max_order_by represents the user_book_reads_max_order_by GraphQL type
type user_book_reads_max_order_by struct {
}

// user_book_reads_min_fields represents the user_book_reads_min_fields GraphQL type
type user_book_reads_min_fields struct {
	Edition_id int `json:"edition_id"`
	Finished_at *date `json:"finished_at"`
	ID int `json:"id"`
	Paused_at *date `json:"paused_at"`
	Progress *float8 `json:"progress"`
	Progress_pages int `json:"progress_pages"`
	Progress_seconds int `json:"progress_seconds"`
	Started_at *date `json:"started_at"`
	User_book_id int `json:"user_book_id"`
}

// user_book_reads_min_order_by represents the user_book_reads_min_order_by GraphQL type
type user_book_reads_min_order_by struct {
}

// user_book_reads_order_by represents the user_book_reads_order_by GraphQL type
type user_book_reads_order_by struct {
}

// user_book_reads_select_column represents the user_book_reads_select_column GraphQL type
type user_book_reads_select_column struct {
}

// user_book_reads_select_column_user_book_reads_aggregate_bool_exp_avg_arguments_columns represents the user_book_reads_select_column_user_book_reads_aggregate_bool_exp_avg_arguments_columns GraphQL type
type user_book_reads_select_column_user_book_reads_aggregate_bool_exp_avg_arguments_columns struct {
}

// user_book_reads_select_column_user_book_reads_aggregate_bool_exp_corr_arguments_columns represents the user_book_reads_select_column_user_book_reads_aggregate_bool_exp_corr_arguments_columns GraphQL type
type user_book_reads_select_column_user_book_reads_aggregate_bool_exp_corr_arguments_columns struct {
}

// user_book_reads_select_column_user_book_reads_aggregate_bool_exp_covar_samp_arguments_columns represents the user_book_reads_select_column_user_book_reads_aggregate_bool_exp_covar_samp_arguments_columns GraphQL type
type user_book_reads_select_column_user_book_reads_aggregate_bool_exp_covar_samp_arguments_columns struct {
}

// user_book_reads_select_column_user_book_reads_aggregate_bool_exp_max_arguments_columns represents the user_book_reads_select_column_user_book_reads_aggregate_bool_exp_max_arguments_columns GraphQL type
type user_book_reads_select_column_user_book_reads_aggregate_bool_exp_max_arguments_columns struct {
}

// user_book_reads_select_column_user_book_reads_aggregate_bool_exp_min_arguments_columns represents the user_book_reads_select_column_user_book_reads_aggregate_bool_exp_min_arguments_columns GraphQL type
type user_book_reads_select_column_user_book_reads_aggregate_bool_exp_min_arguments_columns struct {
}

// user_book_reads_select_column_user_book_reads_aggregate_bool_exp_stddev_samp_arguments_columns represents the user_book_reads_select_column_user_book_reads_aggregate_bool_exp_stddev_samp_arguments_columns GraphQL type
type user_book_reads_select_column_user_book_reads_aggregate_bool_exp_stddev_samp_arguments_columns struct {
}

// user_book_reads_select_column_user_book_reads_aggregate_bool_exp_sum_arguments_columns represents the user_book_reads_select_column_user_book_reads_aggregate_bool_exp_sum_arguments_columns GraphQL type
type user_book_reads_select_column_user_book_reads_aggregate_bool_exp_sum_arguments_columns struct {
}

// user_book_reads_select_column_user_book_reads_aggregate_bool_exp_var_samp_arguments_columns represents the user_book_reads_select_column_user_book_reads_aggregate_bool_exp_var_samp_arguments_columns GraphQL type
type user_book_reads_select_column_user_book_reads_aggregate_bool_exp_var_samp_arguments_columns struct {
}

// user_book_reads_stddev_fields represents the user_book_reads_stddev_fields GraphQL type
type user_book_reads_stddev_fields struct {
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Progress float64 `json:"progress"`
	Progress_pages float64 `json:"progress_pages"`
	Progress_seconds float64 `json:"progress_seconds"`
	User_book_id float64 `json:"user_book_id"`
}

// user_book_reads_stddev_order_by represents the user_book_reads_stddev_order_by GraphQL type
type user_book_reads_stddev_order_by struct {
}

// user_book_reads_stddev_pop_fields represents the user_book_reads_stddev_pop_fields GraphQL type
type user_book_reads_stddev_pop_fields struct {
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Progress float64 `json:"progress"`
	Progress_pages float64 `json:"progress_pages"`
	Progress_seconds float64 `json:"progress_seconds"`
	User_book_id float64 `json:"user_book_id"`
}

// user_book_reads_stddev_pop_order_by represents the user_book_reads_stddev_pop_order_by GraphQL type
type user_book_reads_stddev_pop_order_by struct {
}

// user_book_reads_stddev_samp_fields represents the user_book_reads_stddev_samp_fields GraphQL type
type user_book_reads_stddev_samp_fields struct {
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Progress float64 `json:"progress"`
	Progress_pages float64 `json:"progress_pages"`
	Progress_seconds float64 `json:"progress_seconds"`
	User_book_id float64 `json:"user_book_id"`
}

// user_book_reads_stddev_samp_order_by represents the user_book_reads_stddev_samp_order_by GraphQL type
type user_book_reads_stddev_samp_order_by struct {
}

// user_book_reads_stream_cursor_input represents the user_book_reads_stream_cursor_input GraphQL type
type user_book_reads_stream_cursor_input struct {
}

// user_book_reads_stream_cursor_value_input represents the user_book_reads_stream_cursor_value_input GraphQL type
type user_book_reads_stream_cursor_value_input struct {
}

// user_book_reads_sum_fields represents the user_book_reads_sum_fields GraphQL type
type user_book_reads_sum_fields struct {
	Edition_id int `json:"edition_id"`
	ID int `json:"id"`
	Progress *float8 `json:"progress"`
	Progress_pages int `json:"progress_pages"`
	Progress_seconds int `json:"progress_seconds"`
	User_book_id int `json:"user_book_id"`
}

// user_book_reads_sum_order_by represents the user_book_reads_sum_order_by GraphQL type
type user_book_reads_sum_order_by struct {
}

// user_book_reads_var_pop_fields represents the user_book_reads_var_pop_fields GraphQL type
type user_book_reads_var_pop_fields struct {
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Progress float64 `json:"progress"`
	Progress_pages float64 `json:"progress_pages"`
	Progress_seconds float64 `json:"progress_seconds"`
	User_book_id float64 `json:"user_book_id"`
}

// user_book_reads_var_pop_order_by represents the user_book_reads_var_pop_order_by GraphQL type
type user_book_reads_var_pop_order_by struct {
}

// user_book_reads_var_samp_fields represents the user_book_reads_var_samp_fields GraphQL type
type user_book_reads_var_samp_fields struct {
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Progress float64 `json:"progress"`
	Progress_pages float64 `json:"progress_pages"`
	Progress_seconds float64 `json:"progress_seconds"`
	User_book_id float64 `json:"user_book_id"`
}

// user_book_reads_var_samp_order_by represents the user_book_reads_var_samp_order_by GraphQL type
type user_book_reads_var_samp_order_by struct {
}

// user_book_reads_variance_fields represents the user_book_reads_variance_fields GraphQL type
type user_book_reads_variance_fields struct {
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Progress float64 `json:"progress"`
	Progress_pages float64 `json:"progress_pages"`
	Progress_seconds float64 `json:"progress_seconds"`
	User_book_id float64 `json:"user_book_id"`
}

// user_book_reads_variance_order_by represents the user_book_reads_variance_order_by GraphQL type
type user_book_reads_variance_order_by struct {
}

// user_book_statuses represents the user_book_statuses GraphQL type
type user_book_statuses struct {
	Description string `json:"description"`
	ID int `json:"id"`
	Slug string `json:"slug"`
	Status string `json:"status"`
	User_books []*user_books `json:"user_books"`
	User_books_aggregate *user_books_aggregate `json:"user_books_aggregate"`
}

// user_book_statuses_aggregate represents the user_book_statuses_aggregate GraphQL type
type user_book_statuses_aggregate struct {
	Aggregate *user_book_statuses_aggregate_fields `json:"aggregate"`
	Nodes []*user_book_statuses `json:"nodes"`
}

// user_book_statuses_aggregate_fields represents the user_book_statuses_aggregate_fields GraphQL type
type user_book_statuses_aggregate_fields struct {
	Avg *user_book_statuses_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *user_book_statuses_max_fields `json:"max"`
	Min *user_book_statuses_min_fields `json:"min"`
	Stddev *user_book_statuses_stddev_fields `json:"stddev"`
	Stddev_pop *user_book_statuses_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *user_book_statuses_stddev_samp_fields `json:"stddev_samp"`
	Sum *user_book_statuses_sum_fields `json:"sum"`
	Var_pop *user_book_statuses_var_pop_fields `json:"var_pop"`
	Var_samp *user_book_statuses_var_samp_fields `json:"var_samp"`
	Variance *user_book_statuses_variance_fields `json:"variance"`
}

// user_book_statuses_avg_fields represents the user_book_statuses_avg_fields GraphQL type
type user_book_statuses_avg_fields struct {
	ID float64 `json:"id"`
}

// user_book_statuses_bool_exp represents the user_book_statuses_bool_exp GraphQL type
type user_book_statuses_bool_exp struct {
}

// user_book_statuses_max_fields represents the user_book_statuses_max_fields GraphQL type
type user_book_statuses_max_fields struct {
	Description string `json:"description"`
	ID int `json:"id"`
	Slug string `json:"slug"`
	Status string `json:"status"`
}

// user_book_statuses_min_fields represents the user_book_statuses_min_fields GraphQL type
type user_book_statuses_min_fields struct {
	Description string `json:"description"`
	ID int `json:"id"`
	Slug string `json:"slug"`
	Status string `json:"status"`
}

// user_book_statuses_order_by represents the user_book_statuses_order_by GraphQL type
type user_book_statuses_order_by struct {
}

// user_book_statuses_select_column represents the user_book_statuses_select_column GraphQL type
type user_book_statuses_select_column struct {
}

// user_book_statuses_stddev_fields represents the user_book_statuses_stddev_fields GraphQL type
type user_book_statuses_stddev_fields struct {
	ID float64 `json:"id"`
}

// user_book_statuses_stddev_pop_fields represents the user_book_statuses_stddev_pop_fields GraphQL type
type user_book_statuses_stddev_pop_fields struct {
	ID float64 `json:"id"`
}

// user_book_statuses_stddev_samp_fields represents the user_book_statuses_stddev_samp_fields GraphQL type
type user_book_statuses_stddev_samp_fields struct {
	ID float64 `json:"id"`
}

// user_book_statuses_stream_cursor_input represents the user_book_statuses_stream_cursor_input GraphQL type
type user_book_statuses_stream_cursor_input struct {
}

// user_book_statuses_stream_cursor_value_input represents the user_book_statuses_stream_cursor_value_input GraphQL type
type user_book_statuses_stream_cursor_value_input struct {
}

// user_book_statuses_sum_fields represents the user_book_statuses_sum_fields GraphQL type
type user_book_statuses_sum_fields struct {
	ID int `json:"id"`
}

// user_book_statuses_var_pop_fields represents the user_book_statuses_var_pop_fields GraphQL type
type user_book_statuses_var_pop_fields struct {
	ID float64 `json:"id"`
}

// user_book_statuses_var_samp_fields represents the user_book_statuses_var_samp_fields GraphQL type
type user_book_statuses_var_samp_fields struct {
	ID float64 `json:"id"`
}

// user_book_statuses_variance_fields represents the user_book_statuses_variance_fields GraphQL type
type user_book_statuses_variance_fields struct {
	ID float64 `json:"id"`
}

// user_books represents the user_books GraphQL type
type user_books struct {
	Book *books `json:"book"`
	Book_id int `json:"book_id"`
	Cached_match_score *float8 `json:"cached_match_score"`
	Created_at *timestamptz `json:"created_at"`
	Date_added *date `json:"date_added"`
	Edition *editions `json:"edition"`
	Edition_id int `json:"edition_id"`
	First_read_date *date `json:"first_read_date"`
	First_started_reading_date *date `json:"first_started_reading_date"`
	Followers []*followed_users `json:"followers"`
	Has_review bool `json:"has_review"`
	ID int `json:"id"`
	Imported bool `json:"imported"`
	Last_read_date *date `json:"last_read_date"`
	Likes []*likes `json:"likes"`
	Likes_count int `json:"likes_count"`
	Media_url string `json:"media_url"`
	Merged_at *timestamp `json:"merged_at"`
	Object_type string `json:"object_type"`
	Original_book_id int `json:"original_book_id"`
	Original_edition_id int `json:"original_edition_id"`
	Owned bool `json:"owned"`
	Owned_copies int `json:"owned_copies"`
	Privacy_setting *privacy_settings `json:"privacy_setting"`
	Privacy_setting_id int `json:"privacy_setting_id"`
	Private_notes string `json:"private_notes"`
	Rating *numeric `json:"rating"`
	Read_count int `json:"read_count"`
	Reading_format *reading_formats `json:"reading_format"`
	Reading_format_id int `json:"reading_format_id"`
	Reading_journal_summary *reading_journals_summary `json:"reading_journal_summary"`
	Reading_journals []*reading_journals `json:"reading_journals"`
	Recommended_by string `json:"recommended_by"`
	Recommended_for string `json:"recommended_for"`
	Referrer *users `json:"referrer"`
	Referrer_user_id int `json:"referrer_user_id"`
	Review string `json:"review"`
	Review_has_spoilers bool `json:"review_has_spoilers"`
	Review_html string `json:"review_html"`
	Review_length int `json:"review_length"`
	Review_migrated bool `json:"review_migrated"`
	Review_object *json.RawMessage `json:"review_object"`
	Review_raw string `json:"review_raw"`
	Review_slate *json.RawMessage `json:"review_slate"`
	Reviewed_at *timestamp `json:"reviewed_at"`
	Sponsored_review bool `json:"sponsored_review"`
	Starred bool `json:"starred"`
	Status_id int `json:"status_id"`
	Updated_at *timestamptz `json:"updated_at"`
	URL string `json:"url"`
	User *users `json:"user"`
	User_book_reads []*user_book_reads `json:"user_book_reads"`
	User_book_reads_aggregate *user_book_reads_aggregate `json:"user_book_reads_aggregate"`
	User_book_status *user_book_statuses `json:"user_book_status"`
	User_books []*user_books `json:"user_books"`
	User_books_aggregate *user_books_aggregate `json:"user_books_aggregate"`
	User_id int `json:"user_id"`
}

// user_books_aggregate represents the user_books_aggregate GraphQL type
type user_books_aggregate struct {
	Aggregate *user_books_aggregate_fields `json:"aggregate"`
	Nodes []*user_books `json:"nodes"`
}

// user_books_aggregate_bool_exp represents the user_books_aggregate_bool_exp GraphQL type
type user_books_aggregate_bool_exp struct {
}

// user_books_aggregate_bool_exp_avg represents the user_books_aggregate_bool_exp_avg GraphQL type
type user_books_aggregate_bool_exp_avg struct {
}

// user_books_aggregate_bool_exp_bool_and represents the user_books_aggregate_bool_exp_bool_and GraphQL type
type user_books_aggregate_bool_exp_bool_and struct {
}

// user_books_aggregate_bool_exp_bool_or represents the user_books_aggregate_bool_exp_bool_or GraphQL type
type user_books_aggregate_bool_exp_bool_or struct {
}

// user_books_aggregate_bool_exp_corr represents the user_books_aggregate_bool_exp_corr GraphQL type
type user_books_aggregate_bool_exp_corr struct {
}

// user_books_aggregate_bool_exp_corr_arguments represents the user_books_aggregate_bool_exp_corr_arguments GraphQL type
type user_books_aggregate_bool_exp_corr_arguments struct {
}

// user_books_aggregate_bool_exp_count represents the user_books_aggregate_bool_exp_count GraphQL type
type user_books_aggregate_bool_exp_count struct {
}

// user_books_aggregate_bool_exp_covar_samp represents the user_books_aggregate_bool_exp_covar_samp GraphQL type
type user_books_aggregate_bool_exp_covar_samp struct {
}

// user_books_aggregate_bool_exp_covar_samp_arguments represents the user_books_aggregate_bool_exp_covar_samp_arguments GraphQL type
type user_books_aggregate_bool_exp_covar_samp_arguments struct {
}

// user_books_aggregate_bool_exp_max represents the user_books_aggregate_bool_exp_max GraphQL type
type user_books_aggregate_bool_exp_max struct {
}

// user_books_aggregate_bool_exp_min represents the user_books_aggregate_bool_exp_min GraphQL type
type user_books_aggregate_bool_exp_min struct {
}

// user_books_aggregate_bool_exp_stddev_samp represents the user_books_aggregate_bool_exp_stddev_samp GraphQL type
type user_books_aggregate_bool_exp_stddev_samp struct {
}

// user_books_aggregate_bool_exp_sum represents the user_books_aggregate_bool_exp_sum GraphQL type
type user_books_aggregate_bool_exp_sum struct {
}

// user_books_aggregate_bool_exp_var_samp represents the user_books_aggregate_bool_exp_var_samp GraphQL type
type user_books_aggregate_bool_exp_var_samp struct {
}

// user_books_aggregate_fields represents the user_books_aggregate_fields GraphQL type
type user_books_aggregate_fields struct {
	Avg *user_books_avg_fields `json:"avg"`
	Count int `json:"count"`
	Max *user_books_max_fields `json:"max"`
	Min *user_books_min_fields `json:"min"`
	Stddev *user_books_stddev_fields `json:"stddev"`
	Stddev_pop *user_books_stddev_pop_fields `json:"stddev_pop"`
	Stddev_samp *user_books_stddev_samp_fields `json:"stddev_samp"`
	Sum *user_books_sum_fields `json:"sum"`
	Var_pop *user_books_var_pop_fields `json:"var_pop"`
	Var_samp *user_books_var_samp_fields `json:"var_samp"`
	Variance *user_books_variance_fields `json:"variance"`
}

// user_books_aggregate_order_by represents the user_books_aggregate_order_by GraphQL type
type user_books_aggregate_order_by struct {
}

// user_books_avg_fields represents the user_books_avg_fields GraphQL type
type user_books_avg_fields struct {
	Book_id float64 `json:"book_id"`
	Cached_match_score float64 `json:"cached_match_score"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Owned_copies float64 `json:"owned_copies"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	Rating float64 `json:"rating"`
	Read_count float64 `json:"read_count"`
	Reading_format_id float64 `json:"reading_format_id"`
	Referrer_user_id float64 `json:"referrer_user_id"`
	Review_length float64 `json:"review_length"`
	Status_id float64 `json:"status_id"`
	User_id float64 `json:"user_id"`
}

// user_books_avg_order_by represents the user_books_avg_order_by GraphQL type
type user_books_avg_order_by struct {
}

// user_books_bool_exp represents the user_books_bool_exp GraphQL type
type user_books_bool_exp struct {
}

// user_books_max_fields represents the user_books_max_fields GraphQL type
type user_books_max_fields struct {
	Book_id int `json:"book_id"`
	Cached_match_score *float8 `json:"cached_match_score"`
	Created_at *timestamptz `json:"created_at"`
	Date_added *date `json:"date_added"`
	Edition_id int `json:"edition_id"`
	First_read_date *date `json:"first_read_date"`
	First_started_reading_date *date `json:"first_started_reading_date"`
	ID int `json:"id"`
	Last_read_date *date `json:"last_read_date"`
	Likes_count int `json:"likes_count"`
	Media_url string `json:"media_url"`
	Merged_at *timestamp `json:"merged_at"`
	Object_type string `json:"object_type"`
	Original_book_id int `json:"original_book_id"`
	Original_edition_id int `json:"original_edition_id"`
	Owned_copies int `json:"owned_copies"`
	Privacy_setting_id int `json:"privacy_setting_id"`
	Private_notes string `json:"private_notes"`
	Rating *numeric `json:"rating"`
	Read_count int `json:"read_count"`
	Reading_format_id int `json:"reading_format_id"`
	Recommended_by string `json:"recommended_by"`
	Recommended_for string `json:"recommended_for"`
	Referrer_user_id int `json:"referrer_user_id"`
	Review string `json:"review"`
	Review_html string `json:"review_html"`
	Review_length int `json:"review_length"`
	Review_raw string `json:"review_raw"`
	Reviewed_at *timestamp `json:"reviewed_at"`
	Status_id int `json:"status_id"`
	Updated_at *timestamptz `json:"updated_at"`
	URL string `json:"url"`
	User_id int `json:"user_id"`
}

// user_books_max_order_by represents the user_books_max_order_by GraphQL type
type user_books_max_order_by struct {
}

// user_books_min_fields represents the user_books_min_fields GraphQL type
type user_books_min_fields struct {
	Book_id int `json:"book_id"`
	Cached_match_score *float8 `json:"cached_match_score"`
	Created_at *timestamptz `json:"created_at"`
	Date_added *date `json:"date_added"`
	Edition_id int `json:"edition_id"`
	First_read_date *date `json:"first_read_date"`
	First_started_reading_date *date `json:"first_started_reading_date"`
	ID int `json:"id"`
	Last_read_date *date `json:"last_read_date"`
	Likes_count int `json:"likes_count"`
	Media_url string `json:"media_url"`
	Merged_at *timestamp `json:"merged_at"`
	Object_type string `json:"object_type"`
	Original_book_id int `json:"original_book_id"`
	Original_edition_id int `json:"original_edition_id"`
	Owned_copies int `json:"owned_copies"`
	Privacy_setting_id int `json:"privacy_setting_id"`
	Private_notes string `json:"private_notes"`
	Rating *numeric `json:"rating"`
	Read_count int `json:"read_count"`
	Reading_format_id int `json:"reading_format_id"`
	Recommended_by string `json:"recommended_by"`
	Recommended_for string `json:"recommended_for"`
	Referrer_user_id int `json:"referrer_user_id"`
	Review string `json:"review"`
	Review_html string `json:"review_html"`
	Review_length int `json:"review_length"`
	Review_raw string `json:"review_raw"`
	Reviewed_at *timestamp `json:"reviewed_at"`
	Status_id int `json:"status_id"`
	Updated_at *timestamptz `json:"updated_at"`
	URL string `json:"url"`
	User_id int `json:"user_id"`
}

// user_books_min_order_by represents the user_books_min_order_by GraphQL type
type user_books_min_order_by struct {
}

// user_books_order_by represents the user_books_order_by GraphQL type
type user_books_order_by struct {
}

// user_books_select_column represents the user_books_select_column GraphQL type
type user_books_select_column struct {
}

// user_books_select_column_user_books_aggregate_bool_exp_avg_arguments_columns represents the user_books_select_column_user_books_aggregate_bool_exp_avg_arguments_columns GraphQL type
type user_books_select_column_user_books_aggregate_bool_exp_avg_arguments_columns struct {
}

// user_books_select_column_user_books_aggregate_bool_exp_bool_and_arguments_columns represents the user_books_select_column_user_books_aggregate_bool_exp_bool_and_arguments_columns GraphQL type
type user_books_select_column_user_books_aggregate_bool_exp_bool_and_arguments_columns struct {
}

// user_books_select_column_user_books_aggregate_bool_exp_bool_or_arguments_columns represents the user_books_select_column_user_books_aggregate_bool_exp_bool_or_arguments_columns GraphQL type
type user_books_select_column_user_books_aggregate_bool_exp_bool_or_arguments_columns struct {
}

// user_books_select_column_user_books_aggregate_bool_exp_corr_arguments_columns represents the user_books_select_column_user_books_aggregate_bool_exp_corr_arguments_columns GraphQL type
type user_books_select_column_user_books_aggregate_bool_exp_corr_arguments_columns struct {
}

// user_books_select_column_user_books_aggregate_bool_exp_covar_samp_arguments_columns represents the user_books_select_column_user_books_aggregate_bool_exp_covar_samp_arguments_columns GraphQL type
type user_books_select_column_user_books_aggregate_bool_exp_covar_samp_arguments_columns struct {
}

// user_books_select_column_user_books_aggregate_bool_exp_max_arguments_columns represents the user_books_select_column_user_books_aggregate_bool_exp_max_arguments_columns GraphQL type
type user_books_select_column_user_books_aggregate_bool_exp_max_arguments_columns struct {
}

// user_books_select_column_user_books_aggregate_bool_exp_min_arguments_columns represents the user_books_select_column_user_books_aggregate_bool_exp_min_arguments_columns GraphQL type
type user_books_select_column_user_books_aggregate_bool_exp_min_arguments_columns struct {
}

// user_books_select_column_user_books_aggregate_bool_exp_stddev_samp_arguments_columns represents the user_books_select_column_user_books_aggregate_bool_exp_stddev_samp_arguments_columns GraphQL type
type user_books_select_column_user_books_aggregate_bool_exp_stddev_samp_arguments_columns struct {
}

// user_books_select_column_user_books_aggregate_bool_exp_sum_arguments_columns represents the user_books_select_column_user_books_aggregate_bool_exp_sum_arguments_columns GraphQL type
type user_books_select_column_user_books_aggregate_bool_exp_sum_arguments_columns struct {
}

// user_books_select_column_user_books_aggregate_bool_exp_var_samp_arguments_columns represents the user_books_select_column_user_books_aggregate_bool_exp_var_samp_arguments_columns GraphQL type
type user_books_select_column_user_books_aggregate_bool_exp_var_samp_arguments_columns struct {
}

// user_books_stddev_fields represents the user_books_stddev_fields GraphQL type
type user_books_stddev_fields struct {
	Book_id float64 `json:"book_id"`
	Cached_match_score float64 `json:"cached_match_score"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Owned_copies float64 `json:"owned_copies"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	Rating float64 `json:"rating"`
	Read_count float64 `json:"read_count"`
	Reading_format_id float64 `json:"reading_format_id"`
	Referrer_user_id float64 `json:"referrer_user_id"`
	Review_length float64 `json:"review_length"`
	Status_id float64 `json:"status_id"`
	User_id float64 `json:"user_id"`
}

// user_books_stddev_order_by represents the user_books_stddev_order_by GraphQL type
type user_books_stddev_order_by struct {
}

// user_books_stddev_pop_fields represents the user_books_stddev_pop_fields GraphQL type
type user_books_stddev_pop_fields struct {
	Book_id float64 `json:"book_id"`
	Cached_match_score float64 `json:"cached_match_score"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Owned_copies float64 `json:"owned_copies"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	Rating float64 `json:"rating"`
	Read_count float64 `json:"read_count"`
	Reading_format_id float64 `json:"reading_format_id"`
	Referrer_user_id float64 `json:"referrer_user_id"`
	Review_length float64 `json:"review_length"`
	Status_id float64 `json:"status_id"`
	User_id float64 `json:"user_id"`
}

// user_books_stddev_pop_order_by represents the user_books_stddev_pop_order_by GraphQL type
type user_books_stddev_pop_order_by struct {
}

// user_books_stddev_samp_fields represents the user_books_stddev_samp_fields GraphQL type
type user_books_stddev_samp_fields struct {
	Book_id float64 `json:"book_id"`
	Cached_match_score float64 `json:"cached_match_score"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Owned_copies float64 `json:"owned_copies"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	Rating float64 `json:"rating"`
	Read_count float64 `json:"read_count"`
	Reading_format_id float64 `json:"reading_format_id"`
	Referrer_user_id float64 `json:"referrer_user_id"`
	Review_length float64 `json:"review_length"`
	Status_id float64 `json:"status_id"`
	User_id float64 `json:"user_id"`
}

// user_books_stddev_samp_order_by represents the user_books_stddev_samp_order_by GraphQL type
type user_books_stddev_samp_order_by struct {
}

// user_books_stream_cursor_input represents the user_books_stream_cursor_input GraphQL type
type user_books_stream_cursor_input struct {
}

// user_books_stream_cursor_value_input represents the user_books_stream_cursor_value_input GraphQL type
type user_books_stream_cursor_value_input struct {
}

// user_books_sum_fields represents the user_books_sum_fields GraphQL type
type user_books_sum_fields struct {
	Book_id int `json:"book_id"`
	Cached_match_score *float8 `json:"cached_match_score"`
	Edition_id int `json:"edition_id"`
	ID int `json:"id"`
	Likes_count int `json:"likes_count"`
	Original_book_id int `json:"original_book_id"`
	Original_edition_id int `json:"original_edition_id"`
	Owned_copies int `json:"owned_copies"`
	Privacy_setting_id int `json:"privacy_setting_id"`
	Rating *numeric `json:"rating"`
	Read_count int `json:"read_count"`
	Reading_format_id int `json:"reading_format_id"`
	Referrer_user_id int `json:"referrer_user_id"`
	Review_length int `json:"review_length"`
	Status_id int `json:"status_id"`
	User_id int `json:"user_id"`
}

// user_books_sum_order_by represents the user_books_sum_order_by GraphQL type
type user_books_sum_order_by struct {
}

// user_books_var_pop_fields represents the user_books_var_pop_fields GraphQL type
type user_books_var_pop_fields struct {
	Book_id float64 `json:"book_id"`
	Cached_match_score float64 `json:"cached_match_score"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Owned_copies float64 `json:"owned_copies"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	Rating float64 `json:"rating"`
	Read_count float64 `json:"read_count"`
	Reading_format_id float64 `json:"reading_format_id"`
	Referrer_user_id float64 `json:"referrer_user_id"`
	Review_length float64 `json:"review_length"`
	Status_id float64 `json:"status_id"`
	User_id float64 `json:"user_id"`
}

// user_books_var_pop_order_by represents the user_books_var_pop_order_by GraphQL type
type user_books_var_pop_order_by struct {
}

// user_books_var_samp_fields represents the user_books_var_samp_fields GraphQL type
type user_books_var_samp_fields struct {
	Book_id float64 `json:"book_id"`
	Cached_match_score float64 `json:"cached_match_score"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Owned_copies float64 `json:"owned_copies"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	Rating float64 `json:"rating"`
	Read_count float64 `json:"read_count"`
	Reading_format_id float64 `json:"reading_format_id"`
	Referrer_user_id float64 `json:"referrer_user_id"`
	Review_length float64 `json:"review_length"`
	Status_id float64 `json:"status_id"`
	User_id float64 `json:"user_id"`
}

// user_books_var_samp_order_by represents the user_books_var_samp_order_by GraphQL type
type user_books_var_samp_order_by struct {
}

// user_books_variance_fields represents the user_books_variance_fields GraphQL type
type user_books_variance_fields struct {
	Book_id float64 `json:"book_id"`
	Cached_match_score float64 `json:"cached_match_score"`
	Edition_id float64 `json:"edition_id"`
	ID float64 `json:"id"`
	Likes_count float64 `json:"likes_count"`
	Original_book_id float64 `json:"original_book_id"`
	Original_edition_id float64 `json:"original_edition_id"`
	Owned_copies float64 `json:"owned_copies"`
	Privacy_setting_id float64 `json:"privacy_setting_id"`
	Rating float64 `json:"rating"`
	Read_count float64 `json:"read_count"`
	Reading_format_id float64 `json:"reading_format_id"`
	Referrer_user_id float64 `json:"referrer_user_id"`
	Review_length float64 `json:"review_length"`
	Status_id float64 `json:"status_id"`
	User_id float64 `json:"user_id"`
}

// user_books_variance_order_by represents the user_books_variance_order_by GraphQL type
type user_books_variance_order_by struct {
}

// user_flags represents the user_flags GraphQL type
type user_flags struct {
	Action_id int `json:"action_id"`
	Action_type string `json:"action_type"`
	Category string `json:"category"`
	Created_at *timestamptz `json:"created_at"`
	Details string `json:"details"`
	Flag_status *flag_statuses `json:"flag_status"`
	Flag_status_id int `json:"flag_status_id"`
	ID int `json:"id"`
	Reported_user_id int `json:"reported_user_id"`
	User_id int `json:"user_id"`
	User_reported *users `json:"user_reported"`
	User_submitted *users `json:"user_submitted"`
}

// user_flags_aggregate_order_by represents the user_flags_aggregate_order_by GraphQL type
type user_flags_aggregate_order_by struct {
}

// user_flags_avg_order_by represents the user_flags_avg_order_by GraphQL type
type user_flags_avg_order_by struct {
}

// user_flags_bool_exp represents the user_flags_bool_exp GraphQL type
type user_flags_bool_exp struct {
}

// user_flags_constraint represents the user_flags_constraint GraphQL type
type user_flags_constraint struct {
}

// user_flags_insert_input represents the user_flags_insert_input GraphQL type
type user_flags_insert_input struct {
}

// user_flags_max_order_by represents the user_flags_max_order_by GraphQL type
type user_flags_max_order_by struct {
}

// user_flags_min_order_by represents the user_flags_min_order_by GraphQL type
type user_flags_min_order_by struct {
}

// user_flags_mutation_response represents the user_flags_mutation_response GraphQL type
type user_flags_mutation_response struct {
	Affected_rows int `json:"affected_rows"`
	Returning []*user_flags `json:"returning"`
}

// user_flags_on_conflict represents the user_flags_on_conflict GraphQL type
type user_flags_on_conflict struct {
}

// user_flags_order_by represents the user_flags_order_by GraphQL type
type user_flags_order_by struct {
}

// user_flags_select_column represents the user_flags_select_column GraphQL type
type user_flags_select_column struct {
}

// user_flags_stddev_order_by represents the user_flags_stddev_order_by GraphQL type
type user_flags_stddev_order_by struct {
}

// user_flags_stddev_pop_order_by represents the user_flags_stddev_pop_order_by GraphQL type
type user_flags_stddev_pop_order_by struct {
}

// user_flags_stddev_samp_order_by represents the user_flags_stddev_samp_order_by GraphQL type
type user_flags_stddev_samp_order_by struct {
}

// user_flags_stream_cursor_input represents the user_flags_stream_cursor_input GraphQL type
type user_flags_stream_cursor_input struct {
}

// user_flags_stream_cursor_value_input represents the user_flags_stream_cursor_value_input GraphQL type
type user_flags_stream_cursor_value_input struct {
}

// user_flags_sum_order_by represents the user_flags_sum_order_by GraphQL type
type user_flags_sum_order_by struct {
}

// user_flags_update_column represents the user_flags_update_column GraphQL type
type user_flags_update_column struct {
}

// user_flags_var_pop_order_by represents the user_flags_var_pop_order_by GraphQL type
type user_flags_var_pop_order_by struct {
}

// user_flags_var_samp_order_by represents the user_flags_var_samp_order_by GraphQL type
type user_flags_var_samp_order_by struct {
}

// user_flags_variance_order_by represents the user_flags_variance_order_by GraphQL type
type user_flags_variance_order_by struct {
}

// user_referrals represents the user_referrals GraphQL type
type user_referrals struct {
	Created_at *timestamp `json:"created_at"`
	ID *bigint `json:"id"`
	Referrer *users `json:"referrer"`
	Referrer_id int `json:"referrer_id"`
	State string `json:"state"`
	Updated_at *timestamp `json:"updated_at"`
	User *users `json:"user"`
	User_id int `json:"user_id"`
}

// user_referrals_bool_exp represents the user_referrals_bool_exp GraphQL type
type user_referrals_bool_exp struct {
}

// user_referrals_order_by represents the user_referrals_order_by GraphQL type
type user_referrals_order_by struct {
}

// user_referrals_select_column represents the user_referrals_select_column GraphQL type
type user_referrals_select_column struct {
}

// user_referrals_stream_cursor_input represents the user_referrals_stream_cursor_input GraphQL type
type user_referrals_stream_cursor_input struct {
}

// user_referrals_stream_cursor_value_input represents the user_referrals_stream_cursor_value_input GraphQL type
type user_referrals_stream_cursor_value_input struct {
}

// user_statuses represents the user_statuses GraphQL type
type user_statuses struct {
	ID int `json:"id"`
	Status string `json:"status"`
	Users []*users `json:"users"`
}

// user_statuses_bool_exp represents the user_statuses_bool_exp GraphQL type
type user_statuses_bool_exp struct {
}

// user_statuses_order_by represents the user_statuses_order_by GraphQL type
type user_statuses_order_by struct {
}

// user_statuses_select_column represents the user_statuses_select_column GraphQL type
type user_statuses_select_column struct {
}

// user_statuses_stream_cursor_input represents the user_statuses_stream_cursor_input GraphQL type
type user_statuses_stream_cursor_input struct {
}

// user_statuses_stream_cursor_value_input represents the user_statuses_stream_cursor_value_input GraphQL type
type user_statuses_stream_cursor_value_input struct {
}

// users represents the users GraphQL type
type users struct {
	Access_level int `json:"access_level"`
	Account_privacy_setting_id int `json:"account_privacy_setting_id"`
	Activities []*activities `json:"activities"`
	Activity_privacy_settings_id int `json:"activity_privacy_settings_id"`
	Admin bool `json:"admin"`
	Bio string `json:"bio"`
	Birthdate *date `json:"birthdate"`
	Blocked_users []*user_blocks `json:"blocked_users"`
	Books_count int `json:"books_count"`
	Cached_cover *json.RawMessage `json:"cached_cover"`
	Cached_genres *json.RawMessage `json:"cached_genres"`
	Cached_image *json.RawMessage `json:"cached_image"`
	Collection_imports []*collection_imports `json:"collection_imports"`
	Confirmation_sent_at *timestamp `json:"confirmation_sent_at"`
	Confirmed_at *timestamp `json:"confirmed_at"`
	Created_at *timestamptz `json:"created_at"`
	Current_sign_in_at *timestamp `json:"current_sign_in_at"`
	Email string `json:"email"`
	Email_verified *timestamptz `json:"email_verified"`
	Flair string `json:"flair"`
	Followed_by_users []*followed_users `json:"followed_by_users"`
	Followed_lists []*followed_lists `json:"followed_lists"`
	Followed_prompts []*followed_prompts `json:"followed_prompts"`
	Followed_users []*followed_users `json:"followed_users"`
	Followed_users_count int `json:"followed_users_count"`
	Followers_count int `json:"followers_count"`
	Goals []*goals `json:"goals"`
	ID int `json:"id"`
	Image *images `json:"image"`
	Image_id int `json:"image_id"`
	Last_activity_at *timestamp `json:"last_activity_at"`
	Last_sign_in_at *timestamp `json:"last_sign_in_at"`
	Librarian_roles *json.RawMessage `json:"librarian_roles"`
	Link string `json:"link"`
	Lists []*lists `json:"lists"`
	Lists_aggregate *lists_aggregate `json:"lists_aggregate"`
	Location string `json:"location"`
	Locked_at *timestamp `json:"locked_at"`
	Match_updated_at *timestamp `json:"match_updated_at"`
	Membership string `json:"membership"`
	Membership_ends_at *timestamp `json:"membership_ends_at"`
	Name string `json:"name"`
	Notification_deliveries []*notification_deliveries `json:"notification_deliveries"`
	Notification_deliveries_aggregate *notification_deliveries_aggregate `json:"notification_deliveries_aggregate"`
	Object_type string `json:"object_type"`
	Onboarded bool `json:"onboarded"`
	Payment_system_id int `json:"payment_system_id"`
	Pro bool `json:"pro"`
	Prompt_answers []*prompt_answers `json:"prompt_answers"`
	Prompt_answers_aggregate *prompt_answers_aggregate `json:"prompt_answers_aggregate"`
	Prompts []*prompts `json:"prompts"`
	Pronoun_personal string `json:"pronoun_personal"`
	Pronoun_possessive string `json:"pronoun_possessive"`
	Recommendations []*recommendations `json:"recommendations"`
	Recommended []*recommendations `json:"recommended"`
	Referrer_id int `json:"referrer_id"`
	Referrer_url string `json:"referrer_url"`
	Referrered_users []*user_books `json:"referrered_users"`
	Referrered_users_aggregate *user_books_aggregate `json:"referrered_users_aggregate"`
	Remember_created_at *timestamp `json:"remember_created_at"`
	Reported_user_flags []*user_flags `json:"reported_user_flags"`
	Reset_password_sent_at *timestamp `json:"reset_password_sent_at"`
	Sign_in_count int `json:"sign_in_count"`
	Status_id int `json:"status_id"`
	Taggings []*taggings `json:"taggings"`
	Taggings_aggregate *taggings_aggregate `json:"taggings_aggregate"`
	Unconfirmed_email string `json:"unconfirmed_email"`
	Updated_at *timestamptz `json:"updated_at"`
	User_books []*user_books `json:"user_books"`
	User_books_aggregate *user_books_aggregate `json:"user_books_aggregate"`
	User_flags []*user_flags `json:"user_flags"`
	Username *citext `json:"username"`
}

// users_aggregate_by_created_at_date represents the users_aggregate_by_created_at_date GraphQL type
type users_aggregate_by_created_at_date struct {
	Count *bigint `json:"count"`
	Created_at *date `json:"created_at"`
}

// users_aggregate_by_created_at_date_bool_exp represents the users_aggregate_by_created_at_date_bool_exp GraphQL type
type users_aggregate_by_created_at_date_bool_exp struct {
}

// users_aggregate_by_created_at_date_order_by represents the users_aggregate_by_created_at_date_order_by GraphQL type
type users_aggregate_by_created_at_date_order_by struct {
}

// users_aggregate_by_created_at_date_select_column represents the users_aggregate_by_created_at_date_select_column GraphQL type
type users_aggregate_by_created_at_date_select_column struct {
}

// users_aggregate_by_created_at_date_stream_cursor_input represents the users_aggregate_by_created_at_date_stream_cursor_input GraphQL type
type users_aggregate_by_created_at_date_stream_cursor_input struct {
}

// users_aggregate_by_created_at_date_stream_cursor_value_input represents the users_aggregate_by_created_at_date_stream_cursor_value_input GraphQL type
type users_aggregate_by_created_at_date_stream_cursor_value_input struct {
}

// users_aggregate_order_by represents the users_aggregate_order_by GraphQL type
type users_aggregate_order_by struct {
}

// users_avg_order_by represents the users_avg_order_by GraphQL type
type users_avg_order_by struct {
}

// users_bool_exp represents the users_bool_exp GraphQL type
type users_bool_exp struct {
}

// users_max_order_by represents the users_max_order_by GraphQL type
type users_max_order_by struct {
}

// users_min_order_by represents the users_min_order_by GraphQL type
type users_min_order_by struct {
}

// users_order_by represents the users_order_by GraphQL type
type users_order_by struct {
}

// users_select_column represents the users_select_column GraphQL type
type users_select_column struct {
}

// users_stddev_order_by represents the users_stddev_order_by GraphQL type
type users_stddev_order_by struct {
}

// users_stddev_pop_order_by represents the users_stddev_pop_order_by GraphQL type
type users_stddev_pop_order_by struct {
}

// users_stddev_samp_order_by represents the users_stddev_samp_order_by GraphQL type
type users_stddev_samp_order_by struct {
}

// users_stream_cursor_input represents the users_stream_cursor_input GraphQL type
type users_stream_cursor_input struct {
}

// users_stream_cursor_value_input represents the users_stream_cursor_value_input GraphQL type
type users_stream_cursor_value_input struct {
}

// users_sum_order_by represents the users_sum_order_by GraphQL type
type users_sum_order_by struct {
}

// users_var_pop_order_by represents the users_var_pop_order_by GraphQL type
type users_var_pop_order_by struct {
}

// users_var_samp_order_by represents the users_var_samp_order_by GraphQL type
type users_var_samp_order_by struct {
}

// users_variance_order_by represents the users_variance_order_by GraphQL type
type users_variance_order_by struct {
}

