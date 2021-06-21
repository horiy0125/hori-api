class AddPublishedAtToQiitaPosts < ActiveRecord::Migration[6.1]
  def change
    add_column :qiita_posts, :published_at, :datetime
  end
end
