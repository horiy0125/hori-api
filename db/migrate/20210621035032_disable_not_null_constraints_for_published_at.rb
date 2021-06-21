class DisableNotNullConstraintsForPublishedAt < ActiveRecord::Migration[6.1]
  def change
    change_column_null :announcements, :published_at, true
    change_column_null :qiita_posts, :published_at, true
  end
end
