class CreateAnnouncements < ActiveRecord::Migration[6.1]
  def change
    create_table :announcements do |t|
      t.string :content, null: false
      t.string :link_url
      t.boolean :debug, default: true

      t.timestamps
    end
  end
end
