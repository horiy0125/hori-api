class CreateExternalPosts < ActiveRecord::Migration[6.1]
  def change
    create_table :external_posts do |t|
      t.string :title, null: false
      t.string :url, null: false
      t.string :thumbnail_url, null: false
      t.string :category, null: false
      t.datetime :published_at

      t.timestamps
    end
  end
end
