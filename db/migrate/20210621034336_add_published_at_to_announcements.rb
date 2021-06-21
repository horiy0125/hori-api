class AddPublishedAtToAnnouncements < ActiveRecord::Migration[6.1]
  def change
    add_column :announcements, :published_at, :datetime, null: false
    change_column :qiita_posts, :published_at, :datetime, null: false
  end
end
